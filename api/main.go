package main

import (
	"cadet-project/config"
	"cadet-project/server"
)

func main() {
	config.InitConfig("configurations")
	config.InitDbConfig("configurations")
	server.Run()
}
