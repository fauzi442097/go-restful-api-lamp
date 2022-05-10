package config

import (
	"go-restful-api-lamp/helper"
)

var App = map[string]string{
	"URL": helper.Env("APP_URL", "http://localhost:8080"),
}
