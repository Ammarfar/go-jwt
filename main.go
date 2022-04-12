package main

import (
	c "go-jwt/app/configs"
	h "go-jwt/app/helpers"
	r "go-jwt/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	c.ConnectDatabase()

	app := gin.Default()

	r.Routes(app)

	if err := app.Run(h.GetEnv("APP_SERVER")); err != nil {
		panic(err)
	}
}
