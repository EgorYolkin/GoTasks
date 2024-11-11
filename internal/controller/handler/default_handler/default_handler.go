package default_handler

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/go-co-op/gocron/v2"

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

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	cronTask := func() {
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

	rand.Seed(time.Now().UnixNano())

	j, err := s.NewJob(
		gocron.DurationJob(
			time.Duration(int(time.Hour)*5*rand.Intn(5-2)+2),
		),
		gocron.NewTask(
			cronTask,
		),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(j.ID())

	s.Start()

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      answer,
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		panic(err)
	}
}
