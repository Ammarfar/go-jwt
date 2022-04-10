package routes

import (
	h "go-jwt/app/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, h.ResponseSuccess("welcome to go-jwt app"))
	})

	//Route List
	Auth(app)
	User(app)
}
