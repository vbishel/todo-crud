package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vbishel/trackio-backend/handlers"
)

func main() {
	logger := slog.Default()
	logger.Info("Starting server")

	todoHandler := handlers.NewTodos(logger)

	err := godotenv.Load()
	
	if (err != nil) {
		logger.Error(".env didn't load")
		return
	}

	port, isPortPresent := os.LookupEnv("PORT")
	host, isHostPresent := os.LookupEnv("HOST")

	if !isPortPresent {
		logger.Warn("PORT wasn't found in .env. Using default port 9090")
		port = "9090"
	}

	if !isHostPresent {
		logger.Warn("HOST wasn't found in .env. Using default localhost")
		host = "localhost"
	}

	logger.Info("Server started")

	serveMux := http.NewServeMux()

	serveMux.Handle("/todos", todoHandler)

	s := http.Server{
		Addr: host + ":" + port,
		Handler: serveMux,
	}

	s.ListenAndServe()
}
