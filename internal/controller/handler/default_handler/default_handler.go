package default_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"gotasks/internal/entity"
	"gotasks/internal/usecase/user_usecase"

	"gotasks/internal/repository/storage"
	"gotasks/internal/usecase/data_usecase"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var answer string

	u := entity.User{
		TelegramId: int(update.Message.From.ID),
	}

	stg, _ := ctx.Value("stg").(storage.StorageModel)

	uc := user_usecase.UserUsecase{Storage: stg}
	err := uc.CreateUserIfNotExist(ctx, stg, &u)
	if err != nil {
		panic(err)
	}

	isLink := data_usecase.IsLink(update.Message.Text)

	if !isLink {
		answer = "Кажется, это не ссылка"

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      answer,
			ParseMode: models.ParseModeMarkdown,
		})

		if err != nil {
			panic(err)
		}
		return
	}

	answer = "Продолжаю обработку"

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      answer,
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		panic(err)
	}
}
