package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/doc"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
	"github.com/swaggo/swag"
)

func (s *Server) RegisterDocs() {
	swag.Register(doc.SwaggerInfo.InfoInstanceName, doc.SwaggerInfo)
	s.app.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})
	s.app.Get("/swagger/*", fiberSwagger.HandlerDefault)
	s.app.Get("/openapi.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("doc/swagger.yaml")
	})
}
