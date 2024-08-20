package main

import (
	"lamvng/finance-tracker/controller"
	"lamvng/finance-tracker/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", middleware.CheckAuth, controller.Login)          // Login
	router.POST("/users", controller.CreateUser)                           // Create new user
	router.GET("/accounts", middleware.CheckAuth, controller.ListAccounts) // List all accounts
	router.Run()
}
