package main

import (
	"client/cmd"
	"client/config"
)

func init() {
	config.InitEnvs()
}

func main() {
	cmd.Execute()
}
