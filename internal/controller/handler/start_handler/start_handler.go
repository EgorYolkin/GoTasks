package start_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	answer := "Здравствуй, *" +
		bot.EscapeMarkdown(update.Message.From.FirstName) +
		"*\n\n" +
		"Суть проста — отправляй мне ссылки на статьи/видео, " +
		"а я в случайный момент времени буду отправлять тебе одну из них"
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      answer,
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		panic(err)
	}
}
