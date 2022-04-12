package controllers

import (
	e "go-jwt/app/entities"
	h "go-jwt/app/helpers"
	r "go-jwt/app/repositories"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req e.LoginInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, h.ResponseError("incorrect parameters"))
		return
	}

	user, err := r.FindUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, h.ResponseError(fmt.Sprintf("user %s not found", req.Username)))
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, h.ResponseError("incorrect password"))
		return
	}

	token, err := h.GenerateToken(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, h.ResponseError(err.Error()))
		return
	}

	data := map[string]interface{}{
		"token": token,
	}

	c.JSON(http.StatusOK, h.ResponseSuccess("Successfully Login", data))
}
