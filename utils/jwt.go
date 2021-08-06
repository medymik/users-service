package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(userId int) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorize"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
