package server

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/controllers"
)

var server = controllers.Server{}

func Run() {
	config.InitDbConfig("pkg/config")
	server.InitializeAPI()

	server.Run(config.Config.ApiPort)
}
