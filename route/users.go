package route

import (
	"lamvng/finance-tracker/controller"
	"lamvng/finance-tracker/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController controller.UserControllerInterface) *gin.Engine {

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
