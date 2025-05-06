package main

import (
	"log"
	"os"

	"go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "SERVER STATUS: ", log.LstdFlags)
	srv := server.NewServer(logger)

	logger.Println("Запуст сервера 8080")
	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
