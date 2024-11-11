package data_usecase

import (
	"context"
	"fmt"
	"time"

	"gotasks/internal/entity"
	"gotasks/internal/repository/data_repository"
	"gotasks/internal/repository/storage"
)

type DataUsecase struct {
	Storage storage.StorageModel
}

func (du *DataUsecase) AddData(
	ctx context.Context,
	data entity.Data,
) error {
	repo := data_repository.NewRepository(du.Storage)

	if err := repo.Create(ctx, data); err != nil {
		return err
	}
	return nil
}

func (du *DataUsecase) DeleteData(
	ctx context.Context,
	did uint64,
) error {
	repo := data_repository.NewRepository(du.Storage)

	if err := repo.Delete(ctx, did); err != nil {
		return err
	}

	return nil
}

func (du *DataUsecase) GetRandomData(
	ctx context.Context,
	uid uint64,
) (string, error) {
	repo := data_repository.NewRepository(du.Storage)
	data, err := repo.GetOne(ctx, uid)

	if err != nil {
		return "", err
	}

	t := time.Unix(int64(data.AddedAt), 0)

	answer := fmt.Sprintf(
		"📎 %s\n📖 %s\n🗓️ %s",
		data.Link,
		data.Note,
		t.Format("2 January, 2006"),
	)

	return answer, nil
}
