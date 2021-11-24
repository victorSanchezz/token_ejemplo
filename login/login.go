package login

import (
	"jwt-todo/estructura"
	"jwt-todo/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var User = estructura.User1
var hash = estructura.Hash
var Token string
var err error

func Login(c *gin.Context) {
	var u estructura.User
	hashByte := []byte(hash)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	passwordByte := []byte(u.Password)
	error := bcrypt.CompareHashAndPassword(hashByte, passwordByte)
	if User.Username != u.Username || error != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	Token, err = token.CreateToken(User.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, Token)

}
