package app

import (
	"context"
	"fmt"
	"gotasks/internal/repository/database/postgres"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"

	"gotasks/internal/controller/handler/default_handler"
	"gotasks/internal/controller/handler/get_data_handler"
	"gotasks/internal/controller/handler/start_handler"
	"gotasks/internal/di/di_config"
)

func Run(cfg *di_config.Config) {
	dsn := fmt.Sprintf(
		"user_usecase=%s password=%s host=%s dbname=%s sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.DBName,
	)
	stg, err := postgres.Connect(dsn)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer cancel()

	ctx = context.WithValue(ctx, "stg", stg)

	opts := []bot.Option{
		bot.WithDefaultHandler(default_handler.DefaultHandler),
	}

	b, err := bot.New(cfg.Bot.TelegramApiKey, opts...)

	if err != nil {
		panic(err)
	}

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/start",
		bot.MatchTypeExact,
		start_handler.StartHandler,
	)

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"/get",
		bot.MatchTypeExact,
		get_data_handler.GetDataHandler,
	)

	b.Start(ctx)
}
