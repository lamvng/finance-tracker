package controllers

import (
	"lamvng/finance-tracker/database"
	"lamvng/finance-tracker/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {

	var users models.User
	var newUser models.CreateUserInput

	// Verify input
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify if username exists
	userFound := database.Db.Where("username = ?", newUser.Username).Take(&users)
	if userFound.Error == gorm.ErrRecordNotFound {
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
	user := models.User{
		Username:     newUser.Username,
		PasswordHash: string(passwordHash),
	}
	database.Db.Create(&user)

	// Return OK status
	c.JSON(http.StatusOK, gin.H{"data": user})

}
