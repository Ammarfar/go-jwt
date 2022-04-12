package helpers

import (
	m "go-jwt/app/models"

	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user m.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, m.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		UserID: user.ID,
	})
	tokenString, err := token.SignedString([]byte(GetEnv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}

	authType := strings.Trim(arr[0], "\n\r\t")
	if !(strings.EqualFold(authType, "Bearer")) {
		return "", false
	}

	return strings.Trim(arr[1], "\n\t\r"), true
}

func ValidateToken(tokenString string) (uint, error) {
	var claims m.AuthClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetEnv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	id := claims.UserID
	return id, nil
}
