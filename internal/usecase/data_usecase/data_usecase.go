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

func (du *DataUsecase) GetRandomData(
    ctx context.Context,
    uid uint64,
) (string, error) {
    repo := data_repository.NewDataRepository(du.Storage)
    data, err := repo.GetRandomData(ctx, uid)

    if err != nil {
        return "", err
    }
    
    t := time.Unix(int64(data.AddedAt), 0)
	fmt.Println()
    
    answer := fmt.Sprintf(
        `
            📎 %s\n
            📖 %s\n
            🗓️ %s
        `, 
        data.Link,
        data.Note,
        t.Format("2 January, 2006"),
    )
    
    return answer, nil
}