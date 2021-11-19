package token

import (
	"jwt-todo/estructura"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var User = estructura.User1

func CreateToken(userid uint64) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "")
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["name"] = User.Name
	atClaims["lasname"] = User.Lastname
	atClaims["email"] = User.Email
	atClaims["phone"] = User.Phone
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
