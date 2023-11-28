package middleware

import (
	"main/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey string = "qwerty1234"

func JwtSign(payload models.User) string {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = payload.Id
	atClaims["username"] = payload.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}
