package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

func DefaultEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{ID: "10", Status: "OK"})
	// fmt.Fprintf(w, "api is up and running")
}

func UserTest(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	fmt.Println(username)
}
