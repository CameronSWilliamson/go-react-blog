package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// SQL Config struct
type Config struct {
	ServerName string `json:"server"`
	User       string `json:"username"`
	Password   string `json:"password"`
	DB         string `json:"database"`
}

// Parses a connection string from a config file
var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}

// Reads the config file and returns a Config struct
func readConfig() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Successfully Opened config.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteValue, &config)
	return config
}
