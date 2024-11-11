package default_handler

import (
	"context"
	"fmt"
	"gotasks/internal/usecase/data_usecase"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"golang.org/x/exp/rand"
)

func CronTask(
	ctx context.Context,
	dc data_usecase.DataUsecase,
	update *models.Update,
	b bot.Bot,
) {
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

func ScheduleCron(
	ctx context.Context,
	dc data_usecase.DataUsecase,
	update *models.Update,
	b bot.Bot,
) {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	j, err := s.NewJob(
		gocron.DurationJob(
			time.Duration(int(time.Hour)*5*rand.Intn(5-2)+2),
		),
		gocron.NewTask(
			CronTask,
			ctx,
			dc,
			update,
			b,
		),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(j.ID())

	s.Start()
}
