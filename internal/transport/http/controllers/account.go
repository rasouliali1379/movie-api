package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rasouliali1379/movie-api/internal/config"
	"github.com/rasouliali1379/movie-api/internal/entity/models"
	"github.com/rasouliali1379/movie-api/internal/repository/mongodb"
	"github.com/rasouliali1379/movie-api/internal/service/account"
	"github.com/rasouliali1379/movie-api/internal/transport/http/request_models"
	"github.com/rasouliali1379/movie-api/internal/utils"
	"log"
)

func Login(c *fiber.Ctx) error {
	model := new(request_models.Login)

	if isValid, err := utils.Validate(model, c); !isValid {
		return err
	}

	client, err := mongodb.InitDatabase()

	defer mongodb.DisconnectDatabase(client, context.TODO())

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	service, err := account.NewAccountService(client)

	userId, err := service.Login(model.Email, model.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := utils.GenerateToken(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened, notify developer to fix it",
		})
	}

	return c.JSON(token)
}

func SignUp(c *fiber.Ctx) error {

	model := new(request_models.SignUp)

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(406).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = c.SaveFile(file, fmt.Sprintf("./avatars/%s", file.Filename))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if isValid, err := utils.Validate(model, c); !isValid {
		return err
	}

	client, err := mongodb.InitDatabase()

	defer mongodb.DisconnectDatabase(client, context.TODO())

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	service, err := account.NewAccountService(client)

	if service.EmailExist(model.Email) {
		return c.Status(406).JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	id, err := service.SignUp(models.User{
		FirstName:  model.FirstName,
		LastName:   model.LastName,
		Email:      model.Email,
		Password:   model.Password,
		BirthDate:  model.BirthDate,
		Privileges: model.Privileges,
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

func Refresh(c *fiber.Ctx) error {
	model := new(request_models.Refresh)

	if isValid, err := utils.Validate(model, c); !isValid {
		return err
	}

	token, err := jwt.Parse(model.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		cfg, err := config.GetConfig()

		if err != nil {
			log.Printf("error reading config.yaml file: %s", err)
			return nil, c.Status(500).JSON(fiber.Map{
				"message": "Something unexpected happened. Notify developer to fix it",
			})
		}
		return []byte(cfg.Jwt.Secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid or expired token",
		})
	}

	newToken, err := utils.GenerateToken(claims["id"].(string))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Notify developer to fix it",
		})
	}

	return c.JSON(newToken)
}
