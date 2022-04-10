package configs

import helper "go-jwt/app/helpers"

var JwtKey string = helper.GetEnv("JWT_SECRET_KEY")
var AppServer string = helper.GetEnv("APP_SERVER")
