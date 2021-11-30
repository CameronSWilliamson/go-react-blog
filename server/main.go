package main

import (
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/controllers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/").Subrouter()
	initializeHandlers(subRouter)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/test", controllers.DefaultEndpoint).Methods("GET")
	router.HandleFunc("/users/{user}", controllers.UserTest).Methods("GET")
}
