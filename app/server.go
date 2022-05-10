package app

import (
	"go-restful-api-lamp/config"
	"go-restful-api-lamp/routes"
)

func Run() {

	_ = ConnectDatabase()

	server := routes.Setup()
	server.Run(config.App["url"])

}
