package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/unamdev0/go-todo-app/routes"
)

func main() {
	router := routes.Router()

	//taking port from env file
	port := os.Getenv("PORT")
	fmt.Printf("Starting server at port %s", port)

	//Starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), router))
}
