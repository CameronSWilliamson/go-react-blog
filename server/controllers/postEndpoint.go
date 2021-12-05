package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

func FetchPostsFromUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	posts, err := models.FetchPostsFromUser(mux.Vars(r)["user"])
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(posts)
}
