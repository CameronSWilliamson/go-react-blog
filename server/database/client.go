package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Connector *sql.DB

func Connect() error {
	dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := "root"
	// dbName := "cwilliamson3_DB"

	var err error
	config := readConfig()
	Connector, err = sql.Open(dbDriver, config.User+":"+config.Password+"@("+config.ServerName+")/"+config.DB)
	fmt.Println("Connected to " + config.DB)
	return err
}
