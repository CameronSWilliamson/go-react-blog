package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

func APIEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("{\"resMessage\": \"This is the default endpoint but there's no information\"}")
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func DefaultEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SampleResponse{ID: "10", Status: "OK"})
	// fmt.Fprintf(w, "api is up and running")
}

func UserTest(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["user"]
	fmt.Println(username)
}
