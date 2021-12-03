package models

import "github.com/cameronswilliamson/go-react-blog/database"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Country  string `json:"country"`
}

func FetchAllUsers() ([]User, error) {
	rows, err := database.Connector.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var username string
		var email string
		var country string
		err = rows.Scan(&username, &email, &country)
		if err != nil {
			return nil, err
		}
		user := User{Username: username, Email: email, Country: country}
		users = append(users, user)
	}
	return users, nil
}

func FetchUser(username string) (*User, error) {
	rows, err := database.Connector.Query("SELECT email, country FROM Users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	var email string
	var country string
	for rows.Next() {

		err = rows.Scan(&email, &country)
		if err != nil {
			return nil, err
		}
	}

	user := &User{Username: username, Email: email, Country: country}
	return user, nil
}
