package main

import (
	"cadet-project/configurations"
	"cadet-project/server"
)

func main() {

	configurations.InitConfig("configurations")
	configurations.InitDbConfig("configurations")
	server.Run()
}
