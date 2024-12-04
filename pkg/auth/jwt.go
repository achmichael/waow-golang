package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Token string `json:"token"`
}

func GenerateJWT(username, role string) (*Auth, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24jam
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	auth := &Auth{
		Token: tokenString,
	}

	return auth, nil
}
