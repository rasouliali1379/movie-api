package account

import (
	"errors"
	"github.com/rasouliali1379/movie-api/internal/entity/models"
	"github.com/rasouliali1379/movie-api/internal/repository"
	"github.com/rasouliali1379/movie-api/internal/repository/mongodb"
	"github.com/rasouliali1379/movie-api/internal/service"
	"github.com/rasouliali1379/movie-api/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountService struct {
	db repository.IUserDb
}

func NewAccountService(client *mongo.Client) (service.IAccountService, error) {
	collection, err := mongodb.NewUserDatabase(client)

	if err != nil {
		return nil, err
	}

	return &accountService{
		db: collection,
	}, nil

}

func (service accountService) SignUp(user models.User) (interface{}, error) {

	password := utils.HashAndSalt([]byte(user.Password))
	user.Password = password
	id, err := service.db.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (service accountService) Login(email string, password string) (string, error) {

	user, err := service.db.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	match := utils.ComparePasswords(user.Password, password)

	if !match {
		return "", errors.New("password doesn't match")
	}

	return user.Id.Hex(), nil
}

func (service accountService) EmailExist(email string) bool {
	_, err := service.db.GetUserByEmail(email)

	if err == nil {
		return true
	}
	return false
}
