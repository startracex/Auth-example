package main

import (
	"main/config"
	"main/db"
	"main/router"
)

func main() {
	argp := new(config.Argp)
	argp.New()
	configpath := "config.json"
	argp.StringVar(&configpath, "-c", "--config", "config")
	config.LoadConfig(configpath)
	argp.StringVar(&config.Global.Env, "-e", "--env", "env")
	argp.StringVar(&config.Global.Port, "-p", "--port", "port")

	db.Init()

	router.Router()
}
