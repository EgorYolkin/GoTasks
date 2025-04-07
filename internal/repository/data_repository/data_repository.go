package data_repository

import (
	"context"
	"gotasks/internal/domain"
)

type DataRepository interface {
	Create(ctx context.Context, data domain.Data) error
	GetBy(ctx context.Context, key string, value interface{}) (*domain.Data, error)
	GetAll(ctx context.Context) ([]domain.Data, error)
	DeleteBy(ctx context.Context, key string, value interface{}) error
}
