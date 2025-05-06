package server

import (
	"log"
	"net/http"
	"time"

	"go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func NewServer(logger *log.Logger) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHendler)
	mux.HandleFunc("/upload", handlers.UploaderHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}
