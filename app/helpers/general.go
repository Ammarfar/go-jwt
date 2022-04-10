package helpers

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(name)
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

func GetSession(c *gin.Context) (uint, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return 0, "", false
	}

	username, ok := c.Get("username")
	if !ok {
		return 0, "", false
	}

	return id.(uint), username.(string), true
}
