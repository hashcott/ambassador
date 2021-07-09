package handler

import (
	"github.com/fungerouscode/go-ambassador/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	api := app.Group("api")

	v1 := api.Group("v1")

	admin := v1.Group("auth")

	admin.Post("register", h.Register)
}
