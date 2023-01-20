package server

import (
	"cadet-project/configurations"
	"cadet-project/controllers"
)

var server = controllers.Server{}

func Run() {
	configurations.InitDbConfig("configurations")
	server.InitializeAPI()

	server.Run(configurations.Config.ApiPort)
}
