package main

import (
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/routes"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/gorilla/handlers"
	"github.com/go-openapi/runtime/middleware"
	
)

func main() {
	logs.InfoLogger.Println("Tracified Backend")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	// Start API
	router := routes.NewRouter()

	//serve swagger documentation
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	
	http.Handle("/api/", router)
	logs.InfoLogger.Println("Gateway Started @port " + configs.GetPort() + " with " + configs.EnvName + " environment")
	http.ListenAndServe(configs.GetPort(), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
