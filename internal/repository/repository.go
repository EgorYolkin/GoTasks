package repository

import (
	"context"

	"gotasks/internal/repository/storage"
)

type RepositoryInterface interface {
	NewRepository(*storage.StorageModel) struct{}

	Create(context.Context, struct{}) error
	GetOne(context.Context, uint64) error
	GetAll() ([]struct{}, error)
	Delete(context.Context, uint64) error
}

type Repository struct{}
