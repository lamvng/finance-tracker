package middleware

import (
	"fmt"
	"lamvng/finance-tracker/configs"
	"lamvng/finance-tracker/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	UserService service.UserServiceInterface
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Find authorization token in header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify token format
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify token
		tokenString := authToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(configs.GetEnv("JWT_TOKEN_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Printf("%s\n%t\n", claims, ok)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// if float64(time.Now().Unix()) > claims["exp"].(float64) {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// err := m.UserService.
		// // var user model.User
		// // error := database.DB.Where("ID=?", claims["id"]).Find(&user).Error

		// if error != nil {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		// c.Set("currentUser", user)

		// c.Next()
	}
}
