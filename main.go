package main

import (
	"github.com/Jamsek-m/config-server/config"
	"github.com/Jamsek-m/config-server/db"
	cors "github.com/Jamsek-m/config-server/middlewares"
	"github.com/Jamsek-m/config-server/server"
)

func main() {
	config.ReadConfiguration()
	db.ConnectToDatabase()
	cors.InitializeCorsFilter()
	server.StartServer()
}
