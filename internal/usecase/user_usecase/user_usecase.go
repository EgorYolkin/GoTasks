package user_usecase

import (
	"context"

	"gotasks/internal/entity"
	"gotasks/internal/repository/storage"
	"gotasks/internal/repository/user_repository"
)

type UserUsecase struct {
	Storage storage.StorageModel
}

func (uc *UserUsecase) CreateUserIfNotExist(
	ctx context.Context,
	u *entity.User,
) error {
	repo := user_repository.NewRepository(uc.Storage)

	if err := repo.Create(ctx, *u); err != nil {
		return err
	}

	return nil
}
