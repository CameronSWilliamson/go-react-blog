package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

var Connector *gorm.DB

func Connect(connString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful")
	return nil
}
