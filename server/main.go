package main

import (
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/controllers"
	"github.com/cameronswilliamson/go-react-blog/database"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/").Subrouter()
	initializeHandlers(subRouter)
	initDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/", controllers.APIEndpoint).Methods("GET")
	router.HandleFunc("/test", controllers.DefaultEndpoint).Methods("GET")
	router.HandleFunc("/users/{user}", controllers.FetchSpecificUser).Methods("GET")
	router.HandleFunc("/users/", controllers.FetchUsers).Methods("GET")
}

func initDB() {
	// config := database.Config{ServerName: "db", User: "root", Password: "root", DB: "sys"}
	// cnxstr := database.GetConnectionString(config)
	// err := database.Connect(cnxstr)
	err := database.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
