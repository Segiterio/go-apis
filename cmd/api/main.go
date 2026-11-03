package main

import (
	"example/backend/internal/config"
	"example/backend/internal/db"
	"example/backend/internal/handlers"
	"example/backend/internal/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {

	// load config
	cfg := config.MustLoad()

	DB, err := db.Connect(cfg.DBUrl)
	
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	lh := handlers.NewListingHandler(DB, logger)

	// create new server
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)
	mux.HandleFunc("POST /listings", lh.Create)

	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      middleware.SetRequestUUIDToHeader(mux),
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
