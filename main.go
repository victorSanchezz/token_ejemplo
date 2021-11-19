package main

import (
	"jwt-todo/estructura"
	"jwt-todo/login"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	User   = estructura.User1
)

func main() {
	router.POST("/login", login.Login)
	log.Fatal(router.Run(":8080"))
}
