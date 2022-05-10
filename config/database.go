package config

import "go-restful-api-lamp/helper"

var DefaultConnection = helper.Env("DB_CONNECTION", "mysql")

var Connections = map[string]map[string]string{
	"mysql": {
		"driver":   "mysql",
		"host":     helper.Env("DB_HOST", "127.0.0.1"),
		"port":     helper.Env("DB_PORT", "3306"),
		"database": helper.Env("DB_DATABASE", "forge"),
		"username": helper.Env("DB_USERNAME", "forge"),
		"password": helper.Env("DB_PASSWORD", ""),
	},
	"pgsql": {
		"driver":   "postgres",
		"host":     helper.Env("DB_HOST", "127.0.0.1"),
		"port":     helper.Env("DB_PORT", "'5432"),
		"database": helper.Env("DB_DATABASE", "forge"),
		"username": helper.Env("DB_USERNAME", "forge"),
		"password": helper.Env("DB_PASSWORD", ""),
	},
}
