package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil{
		logrus.Fatalf("Error loading .env file")
	}

	// routes
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pong")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3300" // default port if not specified in .env
	}

	logrus.Printf("Server started on port %s", port)
	
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		logrus.Fatalf("Error starting server: %v", err)
	}
}