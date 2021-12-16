package models

import (
	"fmt"

	"github.com/cameronswilliamson/go-react-blog/database"
)

type Category struct {
	Name string `json:"category_name"`
}

func FetchAllCategories() ([]Category, error) {
	rows, err := database.Connector.Query("SELECT * FROM Categories")
	if err != nil {
		return nil, err
	}
	var name string
	categories := []Category{}
	for rows.Next() {
		err = rows.Scan(&name)
		fmt.Println(name)
		if err != nil {
			return nil, err
		}
		category := Category{Name: name}
		categories = append(categories, category)
	}
	return categories, nil
}
