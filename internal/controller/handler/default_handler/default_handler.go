package default_handler

import (
	"context"
	"fmt"
	"gotasks/pkg/cron"
	"gotasks/pkg/regexp_checks"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"gotasks/internal/domain"
	"gotasks/internal/repository/storage"
	"gotasks/internal/usecase/data_usecase"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var answer string

	isLink := regexp_checks.TextIsLink(update.Message.Text)

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

	data := domain.Data{
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

	go cron.ScheduleCron(ctx, dc, update, *b)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      answer,
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		panic(err)
	}
}
