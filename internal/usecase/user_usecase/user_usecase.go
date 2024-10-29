package user_usecase

import (
	"context"

	"gotasks/internal/entity"
	"gotasks/internal/repository/storage"
	"gotasks/internal/repository/user_repository"
)

func CreateUserIfNotExist(ctx context.Context, stg storage.StorageModel, u *entity.User) error {
	repo := userrepository.NewUserRepository(stg)

	if err := repo.CreateUser(ctx, u); err != nil {
		return err
	}

	return nil
}
