package main

import (
	"lamvng/finance-tracker/controller"
	"lamvng/finance-tracker/database"
	"lamvng/finance-tracker/model"
	"lamvng/finance-tracker/repository"
	"lamvng/finance-tracker/route"
	"lamvng/finance-tracker/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/glog"
)

func init() {

}
func main() {

	// TODO: Load config here, remove "configs" package

	// Init database connections
	glog.Infoln("Initiating database connections...")
	db := database.InitPostGresConnection()
	validate := validator.New()
	// rdb := database.InitRedisConnection()

	// Migrate RDB tables
	db.AutoMigrate(
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

	// Repository
	userRepository := repository.NewUserRepository(db)

	// Service
	userService := service.NewUserService(userRepository, validate)

	// Controller
	userController := controller.NewUserController(userService)

	// Route
	router := gin.Default()
	route.RegisterUserRoutes(router, userController)
	router.Run()

}
