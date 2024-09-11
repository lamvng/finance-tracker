package controller

import (
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/data/response"
	"lamvng/finance-tracker/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type UserControllerInterface interface {
	Auth(c *gin.Context)
	GetUserProfile(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	// Update(c *gin.Context)
	// Delete(c *gin.Context)
}

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &UserController{userService: service}
}

func (ctl *UserController) Auth(c *gin.Context) {
	var authUser request.AuthenticationRequest
	if err := c.ShouldBindBodyWith(&authUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ctl.userService.Auth(authUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   token,
	}
	c.JSON(http.StatusOK, authResponse)
}

func (ctl *UserController) GetUserProfile(c *gin.Context) {
	username, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"user": username,
	})
}

func (ctl *UserController) Create(c *gin.Context) {
	var newUser request.CreateUserRequest
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.userService.Create(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "user created"})
}

func (ctl *UserController) GetByID(c *gin.Context) {
	userIdString := c.Query("id")
	if userIdString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user id parameter provided"})
		return
	}
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad uuid format"})
		return
	}
	userRes, err := ctl.userService.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   userRes,
	})
}
