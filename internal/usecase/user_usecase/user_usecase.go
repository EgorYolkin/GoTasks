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
	repo := user_repository.NewUserRepository(uc.Storage)

	if err := repo.CreateUser(ctx, u); err != nil {
		return err
	}

	return nil
}
