package usecase

import (
	"task/domain"
)

type RegisterUsecase struct {
	UserRepo domain.UserRepository
}

func NewRegisterUsecase(repo domain.UserRepository) *RegisterUsecase {
	return &RegisterUsecase{UserRepo: repo}
}

func (uc *RegisterUsecase) CreateUser(user *domain.User) error {
	return uc.UserRepo.CreateUser(user)
}
