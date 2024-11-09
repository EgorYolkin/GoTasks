package app

import (
    "os"
	"os/signal"
	"context"

	"github.com/go-telegram/bot"
	
	"gotasks/config"
	"gotasks/internal/repository/postgres"
	"gotasks/internal/controller/handler/default_handler"
	"gotasks/internal/controller/handler/get_data_handler"
	"gotasks/internal/controller/handler/start_handler"
)

func Run(cfg config.Config) {
	stg, err := postgres.Connect(cfg.Database.DSN)
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
