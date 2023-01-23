package server

import (
	"cadet-project/config"
	"cadet-project/controllers"
)

var server = controllers.Server{}

func Run() {
	config.InitDbConfig("configurations")
	server.InitializeAPI()

	server.Run(config.Config.ApiPort)
}
