package usecase

import (
	"task/domain"
)

type LoginUsecase struct {
	UserRepo domain.UserRepository
}

func NewLoginUsecase(repo domain.UserRepository) *LoginUsecase {
	return &LoginUsecase{UserRepo: repo}
}

func (uc *LoginUsecase) Login(credentials domain.Credentials) (string, error) {
	return uc.UserRepo.AuthenticateUser(credentials.Username, credentials.Password)
}
func (uc *LoginUsecase) AuthenticateUser(username, password string) (string, error) {
	return uc.UserRepo.AuthenticateUser(username, password)
}

func (uc *LoginUsecase) ValidateToken(tokenString string) (*domain.Claims, error) {
	return uc.UserRepo.ValidateToken(tokenString)
}
