package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

const jwtSecret = "your_jwt_secret" // Use a mesma chave secreta que você usou para gerar tokens

func AuthRequired() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   []byte(jwtSecret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}
