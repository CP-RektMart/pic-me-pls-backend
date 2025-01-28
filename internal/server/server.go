package server

import (
	"log"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	ServerAddr string
}

type Server struct {
	config Config
	App    *fiber.App
	DB     *database.Store
}

func New(config Config, DB *database.Store) *Server {
	return &Server{
		config: config,
		App:    fiber.New(),
		DB:     DB,
	}
}

func (s *Server) Start() {
	log.Printf("server started on %s", s.config.ServerAddr)
	if err := s.App.Listen(s.config.ServerAddr); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
