package controllers

import (
	"fmt"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/database"
	"github.com/cameronswilliamson/go-react-blog/utils"
	"github.com/gorilla/mux"
)

type postsCategories struct {
	Category string `json:"category_name"`
	PostNum  int    `json:"posts"`
	Likes    int    `json:"likes"`
}

type userPosts struct {
	Username string `json:"username"`
	PostNum  int    `json:"posts"`
	Likes    int    `json:"likes"`
}

func FetchFollowersWhoLiked(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post := vars["post"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rows, err := database.Connector.Query("SELECT COUNT(*) FROM Users_Follows uf JOIN Users_Likes ul on uf.`following`=ul.username JOIN Posts p on ul.post_id = p.post_id WHERE ul.post_id = ?", post)
	var count int
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println(count)
	encoder := utils.GetEncoder(&w)
	encoder.Encode(count)
}

func FetchPostsFromCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := database.Connector.Query("SELECT category_name, COUNT(*) AS posts, likes FROM Posts p JOIN Post_Categories pc USING (post_id) JOIN Categories c USING (category_name) JOIN (SELECT COUNT(*) AS likes, category_name FROM Users_Likes JOIN Post_Categories USING (post_id) GROUP BY category_name) q1 USING (category_name) GROUP BY category_name ORDER BY posts DESC, likes DESC;")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var category_name string
	var postCount int
	var likeCount int
	postCategories := []postsCategories{}
	for rows.Next() {
		err = rows.Scan(&category_name, &postCount, &likeCount)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		postCategories = append(postCategories, postsCategories{category_name, postCount, likeCount})
	}

	fmt.Println(postCategories)
	encoder := utils.GetEncoder(&w)
	encoder.Encode(postCategories)
}

func FetchNewActiveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rows, err := database.Connector.Query("SELECT username, posts, comments FROM Users u JOIN (SELECT username, COUNT(*) AS posts FROM Users_Posts up GROUP BY username) q1 USING (username) JOIN (SELECT username, COUNT(*) AS comments FROM Users_Comments uc GROUP BY username) q2 USING (username) WHERE u.join_date > \"2021-06-01\" OR posts > 1 OR comments > 1 ORDER BY posts DESC, comments DESC;")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var username string
	var postCount int
	var likeCount int
	userPostsList := []userPosts{}
	for rows.Next() {
		err = rows.Scan(&username, &postCount, &likeCount)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		userPostsList = append(userPostsList, userPosts{Username: username, PostNum: postCount, Likes: likeCount})
	}
	fmt.Println(userPostsList)
	encoder := utils.GetEncoder(&w)
	encoder.Encode(userPostsList)
}
