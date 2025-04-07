package start_handler

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"gotasks/internal/entity"
	"gotasks/internal/repository/storage"
	"gotasks/internal/usecase/user"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	u := entity.User{
		TelegramId: int(update.Message.From.ID),
	}

	stg, _ := ctx.Value("stg").(storage.StorageModel)

	uc := user.UserUsecase{Storage: stg}
	err := uc.CreateUserIfNotExist(ctx, &u)
	if err != nil {
		panic(err)
	}

	answer := "Здравствуй, *" +
		bot.EscapeMarkdown(update.Message.From.FirstName) +
		"*\n\n" +
		"Суть проста — отправляй мне ссылки на статьи/видео, " +
		"а я в случайный момент времени буду отправлять тебе одну из них"

	go func() {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      answer,
			ParseMode: models.ParseModeMarkdown,
		})

		if err != nil {
			panic(err)
		}
	}()
}
