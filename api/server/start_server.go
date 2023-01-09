package server

import (
	"cadet-project/configurations"
	"cadet-project/controllers"
)

var server = controllers.Server{}

func Run() {
	configurations.InitDbConfig("configurations")
	server.InitializeDB(configurations.ConfigDB.DBDriver, configurations.ConfigDB.DBUser,
		configurations.ConfigDB.DBPassword, configurations.ConfigDB.DBPort, configurations.ConfigDB.DBHost,
		configurations.ConfigDB.DBName)

	server.Run(configurations.Config.ApiPort)
}
