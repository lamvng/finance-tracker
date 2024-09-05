package route

import (
	"lamvng/finance-tracker/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controller.UserController) *gin.Engine {

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome to homepage")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiRouter := router.Group("/api")
	userRouter := apiRouter.Group("/users")
	// tagRouter.GET("/:userId", userController.GetByID)
	userRouter.POST("", userController.Create)
	// tagRouter.PATCH("/:userId", userController.Update)
	// tagRouter.DELETE("/:userId", userController.Delete)

	return router
}
