package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterMiddlewares(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		
		next.ServeHTTP(w, r)
	})
}

/*

	- check is user have permision to view note
	- check is user have permision to view note
	- check is user have permision to view note

	- check is user have permision to view note
	- check is user have permision to delete note
	- check is user have permision to udpate note

*/