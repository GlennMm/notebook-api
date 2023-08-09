package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	migrations "notebook/migrations"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	vi := viper.New()

	vi.SetConfigName("config")
	vi.SetConfigType("yml")
	vi.AddConfigPath(".")

	err := vi.ReadInConfig()
	if err != nil {
		panic("Failed to read the config file.")
	}

	server_config := vi.GetStringMap("server")

	port := server_config["port"]
	host := server_config["host"]
	db_name := ""

	if server_config["env"] == true {
		database_dev := vi.GetStringMap("database_dev")
		db_name = fmt.Sprint(database_dev["name"])
	} else {
		database_prod := vi.GetStringMap("database_prod")
		db_name = fmt.Sprint(database_prod["name"])
	}
	// logging := server_config["logging"]

	db, err := gorm.Open(sqlite.Open(db_name+".sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, " + err.Error())
	}

	fmt.Println("Connected to databse successfully.")

	migrations.RunMigrations(db)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	RegisterRoutes(router, db)
	RegisterMiddlewares(router)

	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprint(host) + ":" + fmt.Sprint(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println("Server running at " + server.Addr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
