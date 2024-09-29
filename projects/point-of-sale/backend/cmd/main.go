package main

import (
	api "backend/delivery/http"
	"backend/internal/app"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("error loading .env file: %s", err)
	}

	mainApp := app.New()
	server := api.New(mainApp)

	r := mux.NewRouter()
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	r.Use(api.LoggingMiddleware)

	server.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler(r))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
