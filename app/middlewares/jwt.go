package middlewares

import (
	h "go-jwt/app/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	token, ok := h.GetToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, h.ResponseError("Invalid Token Format"))
		c.Abort()
		return
	}

	id, err := h.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, h.ResponseError("Invalid Token"))
		c.Abort()
		return
	}

	c.Set("id", id)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}
