package app

import (
	"go-restful-api-lamp/config"
	"go-restful-api-lamp/routes"
)

func Run() {

	db = ConnectDatabase()

	server := routes.Setup(db)
	server.Run(config.App["Url"])

}
