package default_handler

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"gotasks/internal/entity"
	"gotasks/internal/repository/storage"
	"gotasks/internal/usecase/data_usecase"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var answer string

	isLink := data_usecase.IsLink(update.Message.Text)

	if !isLink {
		answer = "I think it's no link"

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

	stg, _ := ctx.Value("stg").(storage.StorageModel)
	dc := data_usecase.DataUsecase{Storage: stg}

	data := entity.Data{
		User:    uint64(update.Message.From.ID),
		Link:    update.Message.Text,
		Note:    "Hi!",
		AddedAt: int(time.Now().Unix()),
	}

	err := dc.AddData(ctx, data)
	if err != nil {
		fmt.Println(err)
		answer = "Processing error"
	} else {
		answer = "Data added"
	}

	go ScheduleCron(ctx, dc, update, *b)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      answer,
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		panic(err)
	}
}
