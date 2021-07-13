package handler

import (
	"github.com/gofiber/fiber/v2"
)

type registerInput struct {
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	PasswordConf string `json:"passwordConf" binding:"required"`
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var input registerInput
	if err := c.BodyParser(&input); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if input.Password != input.PasswordConf {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password do not match",
		})
	}
	user, _ := h.services.CreateUser(
		input.FirstName,
		input.LastName,
		input.Email,
		input.Password)
	return c.JSON(user)
}
