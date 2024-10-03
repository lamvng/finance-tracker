package main

import (
	"lamvng/finance-tracker/controller"
	"lamvng/finance-tracker/database"
	"lamvng/finance-tracker/model"
	"lamvng/finance-tracker/repository"
	"lamvng/finance-tracker/route"
	"lamvng/finance-tracker/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func init() {

	// Load Envs
	if err := godotenv.Load(".env"); err != nil {
		glog.Fatal(err)
	}

	database.InitPostGresConnection()
	// database.InitRedisConnection()

	// Migrate RDB tables
	database.DB.AutoMigrate(
		&model.AccountType{},
		&model.AssetType{},
		&model.AssetUnit{},
		&model.AssetUnitExchangeRate{},
		&model.User{},
		&model.Lender{},
		&model.Account{},
		&model.TransactionType{},
		&model.TransactionCategory{},
		&model.TransactionSubCategory{},
		&model.InvestmentAccount{},
		&model.LiquidSavingAccount{},
		&model.LiquidSpendingAccount{},
		&model.LiquidSpendingBudget{},
		&model.InvestmentAccountPortfolio{},
		&model.InvestmentTransaction{},
		&model.LiquidSpendingTransaction{},
		&model.LendingTransaction{},
		&model.BudgetCategory{},
	)
}

func main() {

	validate := validator.New()

	// Repository
	userRepository := repository.NewUserRepository(database.DB)

	// Service
	userService := service.NewUserService(userRepository, validate)

	// Controller
	userController := controller.NewUserController(userService)

	// Defaut routes
	router := gin.Default()
	// router.GET("", func(context *gin.Context) {
	// 	context.JSON(http.StatusOK, "welcome to homepage")
	// })
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// API routes
	route.RegisterUserRoutes(router, userController)

	router.Run()
}
