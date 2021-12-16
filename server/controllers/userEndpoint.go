package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

// Fetches a list of all users
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	allUsers, err := models.FetchAllUsers()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(allUsers)
}

// Fetches the details for a specific user
func FetchSpecificUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user, err := models.FetchUser(mux.Vars(r)["user"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(user)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(user)
}
