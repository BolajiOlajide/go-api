package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BolajiOlajide/go-api/controllers"
	"github.com/BolajiOlajide/go-api/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fmt.Println("Starting the application...")
	database.InitializeDB()

	router := mux.NewRouter()

	// routes
	router.HandleFunc("/person", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPerson).Methods("GET")

	port := ":" + os.Getenv("PORT")
	http.ListenAndServe(port, router)
}
