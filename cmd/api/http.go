package api

import (
	"github.com/rasouliali1379/movie-api/internal/config"
	"github.com/rasouliali1379/movie-api/internal/transport/http/fiber"
	"log"
)

func StartApi(port string) error {
	app := fiber.New()

	if port == "" {
		cfg, err := config.GetConfig()

		if err != nil {
			log.Printf("error reading config.yaml file: %s\n", err)
			port = getDefaultPort()
		} else {

			if cfg.Api.Port == "" {
				log.Println("port wasn't defined in config.yaml file")
				port = getDefaultPort()
			} else {
				port = cfg.Api.Port
			}
		}
	}

	return app.Start(port)
}

func getDefaultPort() string {
	log.Println("using default port :3000")
	return ":3000"
}
