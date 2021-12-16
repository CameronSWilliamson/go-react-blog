package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cameronswilliamson/go-react-blog/models"
	"github.com/cameronswilliamson/go-react-blog/utils"
	"github.com/gorilla/mux"
)

func FetchCommentsFromPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	postString := mux.Vars(r)["post_id"]
	fmt.Println(postString)
	postID, err := strconv.Atoi(mux.Vars(r)["post_id"])
	if err != nil {
		fmt.Println(err)
		return
	}
	var allComments []models.Comment
	allComments, err = models.FetchCommentsForPost(postID)
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := utils.GetEncoder(&w)
	encoder.Encode(allComments)
}
