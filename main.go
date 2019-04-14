package main

import (
	"github.com/Jamsek-m/config-server/config"
	"github.com/Jamsek-m/config-server/db"
	"github.com/Jamsek-m/config-server/server"
)

func main() {
	config.ReadConfiguration()
	db.ConnectToDatabase()
	server.StartServer()
}
