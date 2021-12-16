package models

import (
	"database/sql"
	"fmt"

	"github.com/cameronswilliamson/go-react-blog/database"
)

type Post struct {
	Username string `json:"username"`
	Post_id  int    `json:"postid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	PostDate string `json:"post-date"`
}

// Fetches all of the posts from a given user
func FetchPostsFromUser(username string, limit int) ([]Post, error) {
	rows, err := database.Connector.Query("SELECT post_id, title, username as author, content, post_date FROM Posts JOIN Users_Posts USING (post_id) JOIN Users USING (username) WHERE username = ? limit ?;", username, limit)
	if err != nil {
		fmt.Println(err.Error())
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
		post := Post{Post_id: postId, Title: title, Username: author, Content: content, PostDate: postDate}
		posts = append(posts, post)
	}
	return posts, nil
}

// Fetches a post given a postID
func FetchPost(postId int) (*Post, error) {
	rows, err := database.Connector.Query("SELECT post_id, title, username, content, post_date FROM Posts JOIN Users_Posts USING (post_id) JOIN Users USING (username) WHERE post_id = ?;", postId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var title string
	var author string
	var content string
	var postDate string
	var post *Post
	for rows.Next() {
		err = rows.Scan(&postId, &title, &author, &content, &postDate)
		if err != nil {
			return nil, err
		}
		post = &Post{Post_id: postId, Title: title, Username: author, Content: content, PostDate: postDate}
	}
	return post, nil
}

// Inserts the provided Post to the database with username as the author.
func CreatePost(post *Post, username string) error {
	_, err := database.Connector.Exec("INSERT INTO Posts (title, content, post_date) VALUES (?, ?, NOW());", post.Title, post.Content)
	if err != nil {
		return err
	}
	var rows *sql.Rows
	rows, err = database.Connector.Query("SELECT post_id FROM Posts WHERE title = ? AND content = ?;", post.Title, post.Content)
	if err != nil {
		return err
	}
	var postId int
	for rows.Next() {
		err = rows.Scan(&postId)
		// fmt.Printf("%d\n", postId)
		if err != nil {
			return err
		}
	}
	fmt.Println(username)
	fmt.Println(postId)
	_, err = database.Connector.Exec("INSERT INTO Users_Posts (username, post_id) VALUES (?, ?);", &username, postId)
	if err != nil {
		return err
	}
	return nil
}
