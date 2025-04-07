package data

import (
	"context"
	"fmt"
	"time"

	"gotasks/internal/entity"
	"gotasks/internal/repository/data_repository"
)

type DataUsecase struct {
	Repo data_repository.DataRepository
}

func (du *DataUsecase) AddData(
	ctx context.Context,
	data entity.Data,
) error {
	if err := du.Repo.Create(ctx, data); err != nil {
		return err
	}
	return nil
}

func (du *DataUsecase) DeleteData(
	ctx context.Context,
	did uint64,
) error {
	if err := du.Repo.Delete(ctx, did); err != nil {
		return err
	}

	return nil
}

func (du *DataUsecase) GetRandomData(
	ctx context.Context,
	uid uint64,
) (string, error) {
	data, err := du.Repo.GetOne(ctx, uid)

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

	err = du.Repo.Delete(ctx, data.ID)

	if err != nil {
		return "", err
	}

	return answer, nil
}
