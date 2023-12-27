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
  return func(c *gin.Context) {
    excludedPaths := map[string]bool{
      "/api/v1/users/auth/login": false,
      "/api/v1/users/auth/register": false,
      "/api/v1/users/auth/refresh": false,
      "/api/v1/users/auth/register/providers": false,
    }

    if _, ok := excludedPaths[c.Request.URL.Path]; ok {
      c.Next()
      return
    }

    authHeader := c.Request.Header.Get("Authorization")

    if authHeader == "" {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
      return
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return []byte(jwtSecret), nil
    })

    if err != nil || !token.Valid {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
      return
    }

    claims := token.Claims.(jwt.MapClaims)

    c.Set("userId", claims["id"])
    c.Set("username", claims["username"])
    c.Set("email", claims["email"])
    c.Set("imageUrl", claims["imageUrl"])

    c.Next()
  }
}
