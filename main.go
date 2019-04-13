package main

import (
	"./config"
	"./db"
	"./server"
)

func main() {
	config.ReadConfiguration()
	db.ConnectToDatabase()
	server.StartServer()
}
