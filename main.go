package main

import (
	"fmt"
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/routes"
	"github.com/gorilla/handlers"
)

func main() {
	fmt.Println("Tracified Backend")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	// Start API
	router := routes.NewRouter()
	http.Handle("/api/", router)
	fmt.Println("Gateway Started @port " + configs.GetPort() + " with " + configs.EnvName + " environment")
	http.ListenAndServe(configs.GetPort(), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
