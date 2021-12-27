package controllers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rasouliali1379/movie-api/internal/entity/models"
	"github.com/rasouliali1379/movie-api/internal/repository/mongodb"
	"github.com/rasouliali1379/movie-api/internal/service/account"
	"github.com/rasouliali1379/movie-api/internal/transport/http/request_models"
	"github.com/rasouliali1379/movie-api/internal/utils"
	"log"
	"time"
)

func Login(c *fiber.Ctx) error {
	model := new(request_models.Login)

	if err := c.BodyParser(model); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validation := validator.New()

	err := validation.Struct(model)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		errs := utils.TranslateError(err, validation)
		return c.Status(400).JSON(map[string]interface{}{
			"message": "Invalid request body",
			"fields":  errs,
		})
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

	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func SignUp(c *fiber.Ctx) error {

	model := new(request_models.User)

	if err := c.BodyParser(model); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	validation := validator.New()

	err := validation.Struct(model)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		errs := utils.TranslateError(err, validation)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"fields":  errs,
		})
	}

	client, err := mongodb.InitDatabase()

	defer mongodb.DisconnectDatabase(client, context.TODO())

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	service, err := account.NewAccountService(client)

	id, err := service.SignUp(models.User{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Password:  model.Password,
		BirthDate: model.BirthDate,
	})

	if err != nil {
		return c.Status(502).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
