package main

import (
	"proxy-app/api/handlers"
	"proxy-app/api/server"
	"proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}