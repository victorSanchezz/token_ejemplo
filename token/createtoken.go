package token

import (
	"jwt-todo/estructura"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var User = estructura.User1

func CreateToken(Userid uint64) (string, error) {
	os.Setenv("ACCESS_SECRET", "ACCESS_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = Userid
	atClaims["name"] = User.Name
	atClaims["lasname"] = User.Lastname
	atClaims["email"] = User.Email
	atClaims["phone"] = User.Phone
	atClaims["authorized"] = true
	atClaims["exp,omitempty"] = time.Now().Add(5 * time.Minute)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	Token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return Token, nil
}
