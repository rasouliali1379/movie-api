package user

import (
	"context"
	"github.com/rasouliali1379/movie-api/cmd/utils"
	"github.com/rasouliali1379/movie-api/internal/entity/models"
	"github.com/rasouliali1379/movie-api/internal/repository/mongodb"
	"github.com/rasouliali1379/movie-api/internal/service/account"
	"log"
)

func Add() error {
	var firstName string = utils.GetUserInput("First name", true, false)
	var lastName string = utils.GetUserInput("Last name", true, false)
	var email string = utils.GetUserInput("email", true, false)
	var password string = utils.GetUserInput("password", true, true)
	var birthDate string = utils.GetUserInput("birth date", true, false)
	client, err := mongodb.InitDatabase()

	defer mongodb.DisconnectDatabase(client, context.TODO())

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	service, err := account.NewAccountService(client)

	if err != nil {
		log.Fatalf("Unable to get collection: %s", err)
	}

	id, err := service.SignUp(models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		BirthDate: birthDate,
	})

	if err != nil {
		log.Fatalf("Unable sign up: %s", err)
	}

	log.Printf("User registered successfully with id %s", id)
	return nil
}

func Login() error {

	var email string = utils.GetUserInput("Email", true, false)
	var password string = utils.GetUserInput("Password", true, true)
	client, err := mongodb.InitDatabase()

	defer mongodb.DisconnectDatabase(client, context.TODO())

	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err)
	}

	service, err := account.NewAccountService(client)

	if err != nil {
		log.Fatalf("Unable to get collection: %s", err)
	}

	token, err := service.Login(email, password)

	if err != nil {
		log.Fatalf("Unable login: %s", err)
	}

	log.Printf("you're logged in successfuly: %s", token)
	return nil
}
