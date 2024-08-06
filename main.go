package main

import (
	"lamvng/finance-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/auth/signup", controllers.CreateUser)
	router.Run()
}
