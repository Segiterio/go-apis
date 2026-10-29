package main

import (
	"example/backend/internal/config"
	"example/backend/internal/db"
	"example/backend/internal/handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	// load config
	cfg := config.MustLoad()

	DB, err := db.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	lh := handlers.NewListingHandler(DB)

	// create new server
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Database connected")
	log.Printf("Server is running at Port %v", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
