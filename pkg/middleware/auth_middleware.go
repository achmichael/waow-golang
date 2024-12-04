package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"waow-go/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		token := strings.Split(bearerToken, " ")[1]

		result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if claims, ok := result.Claims.(jwt.MapClaims); ok {
			ctx.Set("username", claims["username"])
			ctx.Set("role", claims["role"])
		}

		ctx.Next()
	}
}

func MustAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role != models.ROLE_ADMIN {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "this resource is only for admin",
			})
			return
		}
		ctx.Next()
	}
}

func MustUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role != models.ROLE_USER {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "this resource is only for user",
			})
			return
		}
		ctx.Next()
	}
}
