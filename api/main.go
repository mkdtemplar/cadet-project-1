package main

import (
	"cadet-project/pkg/config"
	"cadet-project/pkg/server"
)

func main() {
	config.InitConfig("pkg/config")
	config.InitDbConfig("pkg/config")
	server.Run()
}
