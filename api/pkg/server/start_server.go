package server

import (
	"cadet-project/pkg/config"
)

var server = Server{}

func Run() {
	config.InitDbConfig("pkg/config")
	server.InitializeAPI()

	server.Run(config.Config.ApiPort)
}
