package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/rasouliali1379/movie-api/internal/config"
	"github.com/rasouliali1379/movie-api/internal/transport/http/response_models"
	"log"
	"time"
)

func GenerateToken(userId string) (response_models.Token, error) {
	accessClaims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("error reading config.yaml file: %s", err)
		return response_models.Token{}, err
	}

	access, err := accessToken.SignedString([]byte(cfg.Jwt.Secret))

	if err != nil {
		return response_models.Token{}, err
	}

	refreshClaims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refresh, err := refreshToken.SignedString([]byte(cfg.Jwt.Secret))

	if err != nil {
		return response_models.Token{}, err
	}

	return response_models.Token{
		Access:  access,
		Refresh: refresh,
	}, nil
}
