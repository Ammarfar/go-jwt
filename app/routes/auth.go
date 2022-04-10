package routes

import (
	c "go-jwt/app/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(app *gin.Engine) {
	app.POST("/login", c.Login)
}
