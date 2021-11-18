package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Name:     "jose",
	Lastname: "jose",
	Email:    "jose@gmail.com",
	Phone:    "1234567890",
}

var (
	router = gin.Default()
	hash   string
)

func main() {

	router.POST("/login", Login)
	log.Fatal(router.Run(":9080"))
}
func Login(c *gin.Context) {
	var u User
	hash = "$2a$10$C96TnfOQ56XQsnxHJTkji.XLVKr.rIerZIHxnfeKh5/RIMQvNp6Ve"
	hashByte := []byte(hash)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	passwordByte := []byte(u.Password)
	error := bcrypt.CompareHashAndPassword(hashByte, passwordByte)
	if user.Username != u.Username || error != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
func CreateToken(userid uint64) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", "")
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["name"] = user.Name
	atClaims["lasname"] = user.Lastname
	atClaims["email"] = user.Email
	atClaims["phone"] = user.Phone
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
