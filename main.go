package main

import (
	"main/conf"
	"main/db"
	"main/router"
)

func main() {
	argp := conf.New()
	configpath := "config.json"
	argp.StringVar(&configpath, "-c", "--config", "config")
	conf.LoadConfig(configpath)
	argp.StringVar(&conf.Global.Env, "-e", "--env", "env")
	argp.StringVar(&conf.Global.Port, "-p", "--port", "port")
	// Init database connection after config loaded
	db.Init()
	// Register router
	router.Entrance()
}
