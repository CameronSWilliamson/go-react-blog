package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Connector *sql.DB

func Connect() error {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "cwilliamson3_DB"

	var err error
	Connector, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@(db)/"+dbName)
	fmt.Println("Connected to " + dbName)
	return err
}
