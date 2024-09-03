package controller

import (
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/database"
	"lamvng/finance-tracker/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	var userFound model.User
	var newUser request.CreateUserRequest

	// Verify input
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify if username exists
	if err := database.DB.Where("username = ?", newUser.Username).Take(&userFound).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// Create password salt & hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user on DB
	user := model.User{
		Username:     newUser.Username,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Email:        newUser.Email,
		PasswordHash: string(passwordHash),
	}
	database.DB.Create(&user)

	// Return OK status
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}
