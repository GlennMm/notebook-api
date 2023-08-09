package main

import (
	"notebook/module/user"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	
	user.RegisterUserRoute(router, db)
	
	
}
