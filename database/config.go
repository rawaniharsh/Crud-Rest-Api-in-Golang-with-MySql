package database

import "fmt"

//Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

//"Connection string" refers to the entire DSN (host, user, password, etc).
var GetConnectionString = func(config Config) string {
	//The fmt.Sprintf function in the GO programming language is a function used to return a formatted string.
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}
