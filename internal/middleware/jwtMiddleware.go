package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string = os.Getenv("JWT_SECRET")

func AuthenticateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		excludedPaths := map[string]bool{
			"/api/v1/users/auth/login":              false,
			"/api/v1/users/auth/register":           false,
			"/api/v1/users/auth/refresh":            false,
			"/api/v1/users/auth/register/providers": false,
		}

		if _, ok := excludedPaths[ctx.Request.URL.Path]; ok {
			ctx.Next()
			return
		}

		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		ctx.Set("userId", claims["id"])
		ctx.Set("username", claims["username"])
		ctx.Set("email", claims["email"])
		ctx.Set("imageUrl", claims["imageUrl"])

		ctx.Next()
	}
}
