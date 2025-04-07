package user

import (
	"context"

	"gotasks/internal/entity"
	"gotasks/internal/repository/user_repository"
)

type UserUsecase struct {
	Repo user_repository.UserRepository
}

func (uc *UserUsecase) CreateUserIfNotExist(
	ctx context.Context,
	u *entity.User,
) error {

	if err := uc.Repo.Create(ctx, *u); err != nil {
		return err
	}

	return nil
}
