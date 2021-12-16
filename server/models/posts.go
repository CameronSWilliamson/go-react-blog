package models

import (
	"database/sql"

	"github.com/cameronswilliamson/go-react-blog/database"
)

type Post struct {
	Username string `json:"username"`
	Post_id  int    `json:"postid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	PostDate string `json:"post-date"`
}

func FetchPostsFromUser(username string) ([]Post, error) {
	rows, err := database.Connector.Query("SELECT post_id, title, content, post_date FROM Posts JOIN Users_Posts USING (post_id) JOIN Users USING (username) WHERE username = ?;", username)
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
		post := Post{Post_id: postId, Title: title, Content: content, PostDate: postDate}
		posts = append(posts, post)
	}
	return posts, nil
}

func CreatePost(post *Post) error {
	_, err := database.Connector.Exec("INSERT INTO Posts (title, content, post_date) VALUES (?, ?, ?);", post.Title, post.Content, post.PostDate)
	if err != nil {
		return err
	}
	var rows *sql.Rows
	rows, err = database.Connector.Query("SELECT post_id FROM Posts WHERE title = ? AND content = ? AND post_date = ?;", post.Title, post.Content, post.PostDate)
	if err != nil {
		return err
	}
	var postId int
	for rows.Next() {
		err = rows.Scan(&postId)
		if err != nil {
			return err
		}
	}
	_, err = database.Connector.Exec("INSERT INTO Users_Posts (username, post_id) VALUES (?, ?);", post.Username, postId)
	if err != nil {
		return err
	}
	return nil
}
