package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	allUsers, err := models.FetchAllUsers()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	strArr, err := models.ParseArray(allUsers)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	fmt.Println(strArr)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(allUsers)
}

func FetchSpecificUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	user, err := models.FetchUser(mux.Vars(r)["user"])
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	fmt.Println(user)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(user)
}
