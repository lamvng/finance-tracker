package controller

import (
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// type UserControllerInterface interface {
// 	// GetByID(c *gin.Context)
// 	Create(c *gin.Context)
// 	// Update(c *gin.Context)
// 	// Delete(c *gin.Context)
// }

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) (userController *UserController) {
	return &UserController{userService: service}
}

func (ctl *UserController) Create(c *gin.Context) {
	var newUser request.CreateUserRequest
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.userService.Create(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "user created"})
}
