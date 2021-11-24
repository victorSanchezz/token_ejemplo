package main

import (
	"jwt-todo/login"
	"log"

	"jwt-todo/token2"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login.Login)
	router.POST("/welcome", token2.TokenValid)
	log.Fatal(router.Run(":8080"))
}
