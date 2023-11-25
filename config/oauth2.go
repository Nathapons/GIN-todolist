package config

import (
	"errors"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GetOauth2() (*oauth2.Config, error) {
	envMap, err := godotenv.Read()
	if err != nil {
		return &oauth2.Config{}, errors.New("The project doesn't have .env file")
	}

	clientId := envMap["CLIENT_ID"]
	clientSecret := envMap["CLENT_SECRET"]

	oauth2Config := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/auth/callback",
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email"},
	}
	return oauth2Config, nil
}
