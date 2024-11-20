package service

import (
	"com.wallet/coding/domain/user"
	"com.wallet/coding/dto"
	userRepository "com.wallet/coding/infrastructure/repository"
	"fmt"
)

type IUserService interface {
	Add(user user.User) (int64, error)
	Login(dto *dto.UserLoginDto) error
}

var _ IUserService = (*IUserServiceImpl)(nil)

type IUserServiceImpl struct {
}

func NewUserServiceImpl() *IUserServiceImpl {
	return &IUserServiceImpl{}
}

func (u *IUserServiceImpl) Add(user user.User) (int64, error) {
	var repo = userRepository.UserRepositoryImpl{}
	result, err := repo.Add(user)
	return result, err
}

func (u *IUserServiceImpl) Login(dto *dto.UserLoginDto) error {
	fmt.Println("login")
	return nil
}
