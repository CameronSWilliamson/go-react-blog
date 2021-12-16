package main

import (
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/controllers"
	"github.com/cameronswilliamson/go-react-blog/database"
	"github.com/gorilla/mux"
)

// Main function for the program
func main() {
	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/").Subrouter()
	initializeHandlers(subRouter)
	initDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Initializes all of the routes for the API
func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/", controllers.APIEndpoint).Methods("GET")
	router.HandleFunc("/test", controllers.DefaultEndpoint).Methods("GET")
	router.HandleFunc("/users/{user}", controllers.FetchSpecificUser).Methods("GET")
	router.HandleFunc("/users/", controllers.FetchUsers).Methods("GET")
	router.HandleFunc("/posts/{post_id}", controllers.FetchPost).Methods("GET")
	router.HandleFunc("/users/{user}/posts/{limit}", controllers.FetchPostsFromUser).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("GET")
	router.HandleFunc("/login", controllers.NewUser).Methods("POST")
	router.HandleFunc("/categories", controllers.FetchCategories).Methods("GET")
	router.HandleFunc("/createpost/{username}", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/olap/followers-who-liked/{post}", controllers.FetchFollowersWhoLiked).Methods("GET")
	router.HandleFunc("/olap/category-counts", controllers.FetchPostsFromCategory).Methods("GET")
	router.HandleFunc("/olap/new-active-users", controllers.FetchNewActiveUsers).Methods("GET")
}

// Initializes the database connection
func initDB() {
	err := database.Connect()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
