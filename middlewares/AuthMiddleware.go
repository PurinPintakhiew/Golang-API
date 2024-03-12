package middlewares

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	SecretKey := os.Getenv("SECRET_KEY")

	authorizationHeader := c.Get("Authorization")

	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Authorization header format"})
	}

	token := parts[1]

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token is not found", "token": token})
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token validation failed", "error": err.Error()})
	}

	return c.Next()
}
