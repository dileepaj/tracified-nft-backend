package main

import (
	"net/http"

	"github.com/dileepaj/tracified-nft-backend/configs"
	"github.com/dileepaj/tracified-nft-backend/routes"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logs.InfoLogger.Println("Tracified Backend")
	err := godotenv.Load()
	if err != nil {
		logrus.Println("Info Issue with loading .env file")
		logs.InfoLogger.Println("Info Issue with loading .env1 file")
	}
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	// Start API
	router := routes.NewRouter()

	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml", Path: "api-docs"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/api-docs", sh)

	http.Handle("/api/", router)
	logs.InfoLogger.Println("Gateway Started @port " + configs.GetPort() + " with " + configs.EnvName + " environment")
	http.ListenAndServe(configs.GetPort(), handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
