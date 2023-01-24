package server

import (
	"cadet-project/pkg/config"
)

var server = Server{}

func Run() {
	config.InitDbConfig("pkg/config")
	server.InitializeServer()

	server.Run(config.Config.ApiPort)
}
