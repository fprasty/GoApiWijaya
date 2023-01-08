package middleware

import (
	"github.com/fprasty/GoApiWijaya/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticateUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}

func IsAuthenticateAdmin(c *fiber.Ctx) error {
	cookie := c.Cookies("Admin-jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()

}
