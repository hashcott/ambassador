package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationCookie = "jwt"
	userCtx             = "userId"
)

func (h *Handler) IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if userId, err := h.services.ParserToken(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	} else {
		c.Locals(userCtx, userId)
		return c.Next()
	}
}

func getUserId(c *fiber.Ctx) (uint, error) {
	id := c.Locals(userCtx)
	idInt, ok := id.(uint)
	if !ok {
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{
			"message": "user id invalid type",
		})
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
