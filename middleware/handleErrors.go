package middleware

import (
	"lamvng/finance-tracker/data/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *response.AppError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"description": e.Error()})
			}
		}
	}
}
