package data_usecase

import (
	"context"
	"fmt"
	"gotasks/internal/repository/storage"
	"time"

	"gotasks/internal/domain"
	"gotasks/internal/repository/data_repository"
)

type DataUsecase struct {
	Repo    data_repository.DataRepository
	Storage storage.StorageModel
}

func NewDataUsecase(repo data_repository.DataRepository, storage storage.StorageModel) *DataUsecase {
	return &DataUsecase{Repo: repo, Storage: storage}
}

const (
	AnswerDataFormatStr string = "📎 %s\n📖 %s\n🗓️ %s"
	AnswerTimeFormatStr        = "2 January, 2006"
)

func (du *DataUsecase) AddData(
	ctx context.Context,
	data domain.Data,
) error {
	if err := du.Repo.Create(ctx, data); err != nil {
		return err
	}
	return nil
}

func (du *DataUsecase) GetRandomData(
	ctx context.Context,
	ID uint64,
) (string, error) {
	data, err := du.Repo.GetBy(ctx, "id", ID)

	if err != nil {
		return "", err
	}

	t := time.Unix(int64(data.AddedAt), 0)

	answer := fmt.Sprintf(
		AnswerDataFormatStr,
		data.Link,
		data.Note,
		t.Format(AnswerTimeFormatStr),
	)

	//err = du.Repo.Delete(ctx, data.ID)
	//
	//if err != nil {
	//	return "", err
	//}

	return answer, nil
}
