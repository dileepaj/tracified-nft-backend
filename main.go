package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dileepaj/tracified-nft-backend/routes"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func getPort() string {
	port := os.Getenv("BE_PORT")
	if port != "" {
		return ":" + port
	}
	return ":6080"
}

func main() {
	fmt.Println("Tracified Backend")
	err := godotenv.Load()
	if err != nil {
		logs.ErrorLogger.Println("Error loading .env1 file")
	}
	port := getPort()
	envName := os.Getenv("BRANCH_NAME")

	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Start API
	router := routes.NewRouter()
	http.Handle("/api/", router)
	fmt.Println("Gateway Started @port " + port + " with " + envName + " environment")
	http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
