package service

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

type User struct {
  ID  string
  Username string
  Email  string
  ImageUrl  string
}

func GenerateJwtToken(user User) (string, int64, error) {
  expiresAt := time.Now().Add(time.Minute * 60).Unix()
  claims := jwt.MapClaims{
    "id": user.ID,
    "username": user.Username,
    "email": user.Email,
    "imageUrl": user.ImageUrl,
    "exp": expiresAt,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
  accessToken, err := token.SignedString(jwtSecret)

  if err != nil {
    log.Fatal(err)
    return "", 0, err
  }

  return accessToken, expiresAt, nil

}

func GenerateRefreshToken(user User) (string, int64, error) {
  expiresAt := time.Now().Add(time.Minute * 60 * 24 * 7).Unix()
  claims := jwt.MapClaims{
    "id": user.ID,
    "username": user.Username,
    "email": user.Email,
    "imageUrl": user.ImageUrl,
    "exp": expiresAt,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  refreshToken, err := token.SignedString(jwtSecret)

  if err != nil {
    log.Fatal(err)
    return "", 0, err
  }

  return refreshToken, expiresAt, nil
}
