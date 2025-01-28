package main

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/config"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/server"
)

func main() {
	config := config.Load()
	store := database.New(database.Config{
		PostgresURL: config.PostgresURL,
	})
	server := server.New(server.Config{
		ServerAddr: config.ServerAddr,
	}, store)

	server.Start()
}
