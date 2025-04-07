package usecase

import (
	"gotasks/config"
	"gotasks/internal/repository"
	"gotasks/internal/usecase/data"
	"gotasks/internal/usecase/user"
)

type UseCase struct {
	dataUsecase data.DataUsecase
	userUsecase user.UserUsecase
}

func NewUseCase(cfg *config.Config, r repository.Repository) *UseCase {
	return &UseCase{
		//	Здесь передаю юзкейсы
	}
}
