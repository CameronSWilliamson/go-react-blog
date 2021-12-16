package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

func FetchPostsFromUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	posts, err := models.FetchPostsFromUser(mux.Vars(r)["user"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	post := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(post)
}
