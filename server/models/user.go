package models

import "github.com/cameronswilliamson/go-react-blog/database"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	JoinDate string `json:"joindate"`
	Bio      string `json:"bio"`
}

// Fetches all users from the database
func FetchAllUsers() ([]User, error) {
	rows, err := database.Connector.Query("SELECT username, join_date, bio FROM Users")
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

// Fetches a specific user's information from the database
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

// Create a new user within the database
func CreateUser(u *User) error {
	_, err := database.Connector.Query("INSERT INTO Users (username, email, join_date, bio) VALUES (?, ?, NOW(), ?)", u.Username, u.Email, u.Bio)
	return err
}

// Check if a user exists within the database
func CheckUser(username string) (bool, error) {
	rows, err := database.Connector.Query("SELECT username FROM Users WHERE username = ?", username)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
