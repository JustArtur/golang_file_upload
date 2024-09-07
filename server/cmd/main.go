package main

import (
	"server/config"
	"server/db"
	"server/server"
)

func init() {
	config.InitEnvs()
	db.ConnectToDB()
}

func main() {
	server.RunServer()
}
