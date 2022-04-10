package main

import (
	c "go-jwt/app/configs"
	r "go-jwt/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	c.ConnectDatabase()

	app := gin.Default()

	r.Routes(app)

	if err := app.Run(c.AppServer); err != nil {
		panic(err)
	}
}
