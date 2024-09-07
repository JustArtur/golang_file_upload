package main

import (
	"server/config"
	"server/db"
)

func init() {
	config.InitEnvs()
	db.ConnectToDB()
}

func main() {

}
