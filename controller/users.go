package controller

import (
	"lamvng/finance-tracker/data/request"
	"lamvng/finance-tracker/data/response"
	"lamvng/finance-tracker/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
		c.JSON(http.StatusBadRequest, response.Response{
			Description: err.Error(),
		})
		return
	}
	token, err := ctl.userService.Auth(authUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Description: err.Error(),
		})
		return
	}
	authResponse := response.Response{
		Data: token,
	}
	c.JSON(http.StatusOK, authResponse)
}

func (ctl *UserController) GetUserProfile(c *gin.Context) {
	userId, _ := c.Get("userId")
	c.JSON(200, gin.H{
		"userId": userId,
	})
}

func (ctl *UserController) Create(c *gin.Context) {
	var newUser request.CreateUserRequest
	if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Description: err.Error(),
		})
		return
	}
	err := ctl.userService.Create(newUser)

	if err != nil {
		c.JSON(err.StatusCode, err)
		c.Error(err)
		return
	}

	// BUG: Catch error if email is found
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, response.Response{
			Description: err.Error(),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			Description: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		Description: "user created",
	})
}

func (ctl *UserController) GetByID(c *gin.Context) {
	contextUserId, isExist := c.Get("userId")
	if !isExist {
		c.JSON(http.StatusBadRequest, response.Response{
			Description: "user ID parameter not found",
		})
		return
	}

	// Assert userId type
	userId, ok := contextUserId.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, response.Response{
			Description: "bad user uuid format",
		})
		return
	}
	userRes, err := ctl.userService.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Description: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Data: userRes,
	})
}
