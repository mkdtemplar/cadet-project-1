// Cadet project API:
//
//	title: This an init project for REST API
//	version: 1.0.0
//	contact: Ivan Markovski
//	email: <ivan.markovski@scalefocus.com>
//
// Schemes: http,https
// Host: localhost:8080
// BasePath: /
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger: meta
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
