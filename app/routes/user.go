package routes

import (
	c "go-jwt/app/controllers"
	m "go-jwt/app/middlewares"

	"github.com/gin-gonic/gin"
)

func User(app *gin.Engine) {
	userRoute := app.Group("/user", m.VerifyToken)
	{
		userRoute.GET("/info", c.GetUserInfo)
	}
}
