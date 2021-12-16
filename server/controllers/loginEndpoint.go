package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
)

// Handles a /login get request to authenticate
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userName := r.Header.Get("username")
	exists, err := models.CheckUser(userName)
	if err != nil || !exists {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if exists {
		fmt.Println(exists)
		w.WriteHeader(http.StatusOK)
	}
	fmt.Println(userName)
}

// Handles a /login post request to create a new user
func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *models.User
	fmt.Println("new user")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user)
	err = models.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
