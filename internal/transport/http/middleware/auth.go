package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/rasouliali1379/movie-api/internal/config"
	"log"
)

func Protected() fiber.Handler {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("error reading config.yaml file: %s", err)
		return nil
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(cfg.Jwt.Secret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"message": "Invalid or expired JWT", "data": nil})
}
