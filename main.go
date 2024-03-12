package main

import (
	"task-list/config"
	"task-list/routers"
)

func main() {
	config.Init()

	r := routers.SetupRouter()
	r.Run(":"+ config.Conf.GetString("PORT"))
}