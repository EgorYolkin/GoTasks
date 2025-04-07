package di_config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Bot struct {
		TelegramApiKey string
	}
	Database struct {
		User     string
		Password string
		Host     string
		DBName   string
	}
}

func NewConfig(envFilePath string) (Config, error) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config

	config.Bot.TelegramApiKey = os.Getenv("TELEGRAM_API_KEY")

	config.Database.User = os.Getenv("DB_USER")
	config.Database.Password = os.Getenv("DB_PASS")
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.DBName = os.Getenv("DB_NAME")

	if config.Bot.TelegramApiKey == "" || config.Database.User == "" || config.Database.Password == "" {
		log.Fatal("Some required environment variables are missing")
		return Config{}, err
	}

	return config, nil
}
