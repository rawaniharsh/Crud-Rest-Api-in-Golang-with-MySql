package database

import (
	"log"
	"rest-go-demo/entity"

	"github.com/jinzhu/gorm"
)

//gorm is an ORM(Object-relational mapping) library for golang
//Object-Relational Mapping (ORM) is a technique of accessing a relational database from an object-oriented language.

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Database connection established!")
	return nil
}

//migrate database table
func Migrate(table *entity.Student) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated successfully")
}
