package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tomascaceres14/go_rest_websockets/handlers"
	"github.com/tomascaceres14/go_rest_websockets/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	config := &server.Config{
		Port:         PORT,
		JWTSecretKey: JWT_SECRET_KEY,
		DatabaseUrl:  DATABASE_URL,
	}

	server, err := server.NewServer(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	server.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
