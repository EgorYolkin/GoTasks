package usecase

import (
	"gotasks/config"
	"gotasks/internal/repository"
)


type UseCase struct {
    
}


func NewUseCase(cfg *config.Config, r repository.Repository) *UseCase {
    return &UseCase{}
}