package main

import (
	"fmt"
	"go-sqlite/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv" // import the godotenv package
)

func init() {
	// Load env variables using the godotenv package
	// DB details is saved as env variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	// application port
	port := os.Getenv("PORT")

	// if port is not set use default port
	if port == "" {
		port = "8000"
	}

	r := router.Router()
	fmt.Println("Starting server on the port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
