package controllers

import (
	h "go-jwt/app/helpers"
	r "go-jwt/app/repositories"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id, _, ok := h.GetSession(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, h.ResponseError("User not found"))
		return
	}

	user, err := r.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, h.ResponseError("User not found"))
		return
	}

	c.JSON(http.StatusOK, h.ResponseError("Success retrieving data!", user))
}
