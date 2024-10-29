package main

import (
	"os"

	"gotasks/config"
	"gotasks/internal/app"
)

func main() {
	cfg := config.Config{
		TelegramApiKey: os.Getenv("TELEGRAM_API_KEY"),
		DSN:            os.Getenv("DSN"),
	}
	app.Run(&cfg)
}
