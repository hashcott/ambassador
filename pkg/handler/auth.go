package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type registerInput struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	PasswordConf string `json:"password_conf" binding:"required"`
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

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var input loginInput
	if err := c.BodyParser(&input); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if token, err := h.services.GenerateToken(input.Email, input.Password); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	} else {
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 12),
			HTTPOnly: true,
		}
		c.Cookie(&cookie)
		return c.JSON(token)
	}
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if userId, err := h.services.ParserToken(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	} else {
		if user, err := h.services.GetUserById(userId); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		} else {
			return c.JSON(user)
		}
	}
}
