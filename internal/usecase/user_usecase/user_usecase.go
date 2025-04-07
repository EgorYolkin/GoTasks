package user_usecase

import (
	"context"
	"gotasks/internal/repository/storage"

	"gotasks/internal/domain"
	"gotasks/internal/repository/user_repository"
)

type UserUsecase struct {
	Repo    user_repository.UserRepository
	Storage storage.StorageModel
}

func NewUserUsecase(repo user_repository.UserRepository, storage storage.StorageModel) *UserUsecase {
	return &UserUsecase{Repo: repo, Storage: storage}
}

func (uc *UserUsecase) CreateUserIfNotExist(
	ctx context.Context,
	u *domain.User,
) error {

	if err := uc.Repo.Create(ctx, *u); err != nil {
		return err
	}

	return nil
}
