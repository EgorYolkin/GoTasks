package main

import (
	"gotasks/cmd/app"
	"gotasks/internal/di/di_config"
	"os"
)

func main() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		panic("ENV_FILE_PATH environment variable not set")
	}

	cfg, err := di_config.InitializeConfig(envFilePath)
	if err != nil {
		panic(err)
	}
	app.Run(cfg)
}
