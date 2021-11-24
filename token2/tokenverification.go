package token2

import (
	"jwt-todo/estructura"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("ACCESS_SECRET")
var User = estructura.User1

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func TokenValid(c *gin.Context) {
	tooken := ExtractToken(c.Request)
	//tooken := login.Token
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tooken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "Invalid Token")
			return
		}
		c.JSON(http.StatusBadRequest, "Invalid Token")
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "Invalid Token")
		return
	}
	c.JSON(http.StatusOK, "Hi!!! "+User.Name)
}
