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
	initializeHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/api/test", controllers.DefaultEndpoint).Methods("GET")
}
