package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

/*
1. create a database crud_restapi
2. create a folder database and file config.go to setup the database config and create a struct for managing a config
3. create a client to manage and create connections to our Database
4. check whether our client can connect with our database in main.go
5. create student struct which would be our object in reference.
6. Create a Simple HTTP Server using gorilla/mux
7. adding a route “/create/” to our router to serve an endpoint which will create a new student whenever triggered for post endpoint
8. create a respective createStudent(), the function that will take the POST request data and create a new person entry in DB.
We will do the same by first Unmarsahlling the JSON data retrieved from the body into our Person struct created above and later insert the data by creating a new record.
*/

func main() {
	initDB()

	//6th step
	log.Println("Http server started on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

//4th step
func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "pass",
			DB:         "crud_restapi",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Student{})
}
