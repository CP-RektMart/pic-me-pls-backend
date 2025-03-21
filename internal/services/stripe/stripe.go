package stripe

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	SecretKey     string `env:"SECRET_KEY"`
	WebhookSecret string `env:"WEBHOOK_SECRET"`
}

type Handler struct {
	store        *database.Store
	validate     *validator.Validate
	stripeConfig Config
	frontendUrl  string
}

func NewHandler(store *database.Store, validate *validator.Validate, stripeConfig Config, frontendUrl string) *Handler {
	return &Handler{
		store:        store,
		validate:     validate,
		stripeConfig: stripeConfig,
		frontendUrl:  frontendUrl,
	}
}
