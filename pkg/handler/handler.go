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

	// Admin router
	admin := v1.Group("auth")
	admin.Post("register", h.Register)
	admin.Post("login", h.Login)

	// Authenticated
	adminAuthenticated := admin.Use(h.IsAuthenticated)
	adminAuthenticated.Get("user", h.GetUser)
	adminAuthenticated.Get("logout", h.Logout)
	adminAuthenticated.Put("user/info", h.UpdateInfo)

}
