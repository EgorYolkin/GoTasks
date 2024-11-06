package get_data_handler

import (
	"context"

	"gotasks/internal/repository/storage"
	"gotasks/internal/usecase/data_usecase"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func GetDataHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	stg, _ := ctx.Value("stg").(storage.StorageModel)
	dc := data_usecase.DataUsecase{Storage: stg}

	answer, err := dc.GetRandomData(ctx, uint64(update.Message.From.ID))

	if len(answer) == 0 {
		answer = "Кажется, данных нет"
	}
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   answer,
	})

	if err != nil {
		panic(err)
	}
}
