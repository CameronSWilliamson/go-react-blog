package models

import "github.com/cameronswilliamson/go-react-blog/database"

type Comment struct {
	CommentID int    `json:"commentid"`
	PostID    int    `json:"postid"`
	Username  string `json:"username"`
	Content   string `json:"content"`
}

func FetchCommentsForPost(postID int) ([]Comment, error) {
	rows, err := database.Connector.Query("SELECT comment_id, content, username FROM Comments JOIN Users_Comments uc USING (comment_id) WHERE post_id = ? ORDER BY comment_date", postID)
	if err != nil {
		return nil, err
	}
	var commentid int
	var content string
	var username string
	comments := []Comment{}

	for rows.Next() {
		err = rows.Scan(&commentid, &content, &username)
		if err != nil {
			return nil, err
		}
		comment := Comment{CommentID: commentid, Content: content, Username: username}
		comments = append(comments, comment)
	}
	return comments, nil
}
