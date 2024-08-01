package main

import (
	"fmt"
	"go-todo-app/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	router := routes.Router()

	//taking port from env file
	port := os.Getenv("PORT")
	fmt.Printf("Starting server at port %s", port)

	//Starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), router))
}
