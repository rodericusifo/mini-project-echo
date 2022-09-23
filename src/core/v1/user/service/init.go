package service_user_v1

import (
	repository_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/repository"
	dto_service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service/dto"
)

type IUserService interface {
	CreateUser(payload *dto_service_user_v1.CreateUserDTO) error
	GetUser(payload *dto_service_user_v1.GetUserDTO) (*dto_service_user_v1.UserDTO, error)
}

type UserService struct {
	UserRepository repository_user_v1.IUserRepository
}

func InitUserService(userRepository repository_user_v1.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}
