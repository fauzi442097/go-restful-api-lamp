package config

import (
	"go-restful-api-lamp/helper"
	"os"
)

var Secret = os.Getenv("JWT_SECRET")
var JwtExpiredAt = helper.Env("JWT_EXPIRED", "60") // default 60 minutes
var Issuer = helper.Env("APP_NAME", "Golang")
