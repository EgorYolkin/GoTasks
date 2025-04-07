package user_repository

import (
	"context"
	"gotasks/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetAll(ctx context.Context) ([]domain.User, error)
}
