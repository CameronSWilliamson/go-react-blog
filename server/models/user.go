package models

import "github.com/cameronswilliamson/go-react-blog/database"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	JoinDate string `json:"joindate"`
	Bio      string `json:"bio"`
}

func FetchAllUsers() ([]User, error) {
	rows, err := database.Connector.Query("SELECT username, joinDate, bio FROM Users")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var username string
		var joinDate string
		var bio string
		err = rows.Scan(&username, &joinDate, &bio)
		if err != nil {
			return nil, err
		}
		user := User{Username: username, JoinDate: joinDate, Bio: bio}
		users = append(users, user)
	}
	return users, nil
}

func FetchUser(username string) (*User, error) {
	rows, err := database.Connector.Query("SELECT join_date, bio FROM Users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	var joinDate string
	var bio string
	for rows.Next() {

		err = rows.Scan(&joinDate, &bio)
		if err != nil {
			return nil, err
		}
	}

	user := &User{Username: username, JoinDate: joinDate, Bio: bio}
	return user, nil
}

func CreateUser(u *User) error {
	_, err := database.Connector.Query("INSERT INTO Users (username, email, join_date, bio) VALUES (?, ?, NOW(), ?)", u.Username, u.Email, u.Bio)
	return err
}
