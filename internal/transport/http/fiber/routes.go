package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasouliali1379/movie-api/internal/transport/http/controllers"
)

func (r *rest) routing() {
	r.fiber.Post("/login", controllers.Login)
	r.fiber.Post("/signup", controllers.SignUp)
	r.fiber.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(map[string]string{
			"message": "Path not found",
		})
	})
}
