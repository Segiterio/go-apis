package main

import (
	"example/backend/internal/config"
	"example/backend/internal/handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	// load config
	cfg := config.MustLoad()

	// create new server
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is running at Port %v", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
