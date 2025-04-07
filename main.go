package main

import (
	"gotasks/cmd/app"
	"gotasks/config"
)

func main() {
	cfg, err := config.NewConfig("configuration.yaml")
	if err != nil {
		panic(err)
	}
	app.Run(cfg)
}
