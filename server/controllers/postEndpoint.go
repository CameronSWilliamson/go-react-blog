package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/gorilla/mux"
)

func FetchPostsFromUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var posts []models.Post
	limit, err := strconv.Atoi(mux.Vars(r)["limit"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	posts, err = models.FetchPostsFromUser(mux.Vars(r)["user"], limit)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(posts)
}

func FetchPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var post *models.Post
	post_id, err := strconv.Atoi(mux.Vars(r)["post_id"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	post, err = models.FetchPost(post_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	username := mux.Vars(r)["username"]
	post := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(post)
	fmt.Println(username)
	err = models.CreatePost(&post, username)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
