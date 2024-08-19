package main

import (
	"lamvng/finance-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", controllers.Login)      // Login
	router.POST("/users", controllers.CreateUser) // Create new user
	router.Run()
}
