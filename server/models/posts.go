package models

import (
	"github.com/cameronswilliamson/go-react-blog/database"
)

type Post struct {
	Post_id  int    `json:"postid"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	PostDate string `json:"post-date"`
}

func FetchPostsFromUser(username string) ([]Post, error) {
	rows, err := database.Connector.Query("SELECT post_id, title, author, content, post_date FROM Posts JOIN Users_Posts USING (post_id) JOIN Users USING (username) WHERE username = ?;", username)
	if err != nil {
		return nil, err
	}
	var postId int
	var title string
	var author string
	var content string
	var postDate string
	posts := []Post{}
	for rows.Next() {
		err = rows.Scan(&postId, &title, &author, &content, &postDate)
		if err != nil {
			return nil, err
		}
		post := Post{Post_id: postId, Title: title, Author: author, Content: content, PostDate: postDate}
		posts = append(posts, post)
	}
	return posts, nil
}
