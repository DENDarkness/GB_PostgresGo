package main

import (
	"app/config"
	"app/server"

	"log"

	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("server.port")); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
