package main

import (
	"os"

	"gotasks/config"
	"gotasks/internal/app"
)

func main() {
	cfg := config.Config{
		TelegramApiKey: os.Getenv("TELEGRAM_API_KEY"),
		DSN:            "user=root password=4806 host=localhost dbname=home sslmode=disable",
	}
	app.Run(&cfg)
}
