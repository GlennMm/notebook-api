package main

import (
	"github.com/gorilla/mux"
	"notebook/module/user"
)

func RegisterRoutes(router *mux.Router) {
	
	user.RegisterUserRoute(router)
	
	
}
