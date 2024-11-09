package main

import (
	"gotasks/config"
	"gotasks/internal/app"
)

func main() {
	cfg, err := config.NewConfig("configuration.yaml")
	if err != nil {
	   panic(err)
	}
	app.Run(cfg)
}
