package app

import (
	"go-restful-api-lamp/config"
	"go-restful-api-lamp/routes"

	"github.com/go-playground/validator/v10"
)

func Run() {

	db = ConnectDatabase()
	validator := validator.New()

	server := routes.Setup(db, validator)
	server.Run(config.App["Url"])

}
