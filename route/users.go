package route

import (
	"lamvng/finance-tracker/controller"
	"lamvng/finance-tracker/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController controller.UserControllerInterface) *gin.Engine {

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome to homepage")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiRouter := router.Group("/api")

	// Login
	loginRouter := apiRouter.Group("/login")
	loginRouter.POST("", userController.Auth)

	// Users
	userRouter := apiRouter.Group("/users")
	userRouter.POST("", userController.Create)
	userRouter.Use(middleware.TokenAuthMiddleware())
	{
		userRouter.GET("", userController.GetByID)
	}

	return router
}
