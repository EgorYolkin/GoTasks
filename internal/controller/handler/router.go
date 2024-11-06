package handler

import (
	"context"
	"gotasks/internal/controller/handler/default_handler"
	"gotasks/internal/controller/handler/start_handler"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Router struct {
	DefaultHandler func(ctx context.Context)
	StartHandler   func(ctx context.Context)
}

func NewRouter(b *bot.Bot, update *models.Update) *Router {
	return &Router{
		StartHandler: func(ctx context.Context) {
			start_handler.StartHandler(ctx, b, update)
		},
		DefaultHandler: func(ctx context.Context) {
			default_handler.DefaultHandler(ctx, b, update)
		},
	}
}
