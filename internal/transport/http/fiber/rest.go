package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasouliali1379/movie-api/internal/transport/http"
)

type rest struct {
	fiber *fiber.App
}

func New() http.IRest {
	return &rest{
		fiber: fiber.New(),
	}
}

func (r *rest) Start(address string) error {
	r.routing()
	return r.fiber.Listen(address)
}

func (r *rest) Shutdown() error {
	//TODO implement me
	panic("implement me")
}
