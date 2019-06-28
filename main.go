package main

import (
	"net/http"
	"jokeapp/backend"
	"github.com/gorilla/handlers"
	"log"
)

//entrypoint
func main() {

	// if changing port change in homepage to serve the ui also(optional)
	port := "5000"
	router := backend.NewRouter() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})
	//TODO : Add middleware for throtling control
	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
