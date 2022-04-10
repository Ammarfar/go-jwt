package middlewares

import (
	c "go-jwt/app/configs"
	h "go-jwt/app/helpers"
	m "go-jwt/app/models"

	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user m.User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, m.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Username,
			ExpiresAt: expiresAt,
		},
		UserID: user.ID,
	})
	tokenString, err := token.SignedString([]byte(c.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (uint, string, error) {
	var claims m.AuthClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.JwtKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	id := claims.UserID
	username := claims.Subject
	return id, username, nil
}

func VerifyToken(c *gin.Context) {
	token, ok := h.GetToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, h.ResponseError("Invalid Token Format"))
		return
	}

	id, username, err := ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, h.ResponseError("Invalid Token"))
		return
	}

	c.Set("id", id)
	c.Set("username", username)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}
