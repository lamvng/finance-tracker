package controllers

import (
	"lamvng/finance-tracker/database"
	"lamvng/finance-tracker/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var authInput models.AuthenticationInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Username not found
	var userFound models.User
	err := database.DB.Where("username=?", authInput.Username).Find(&userFound).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password not correct"})
		return
	}

	// Password not correct
	if err := bcrypt.CompareHashAndPassword([]byte(userFound.PasswordHash), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password not correct"})
		return
	}

	// Create and send login token
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userFound.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}
	glog.Errorf("Failed to generate token: %s\n", err)

	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}
