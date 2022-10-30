package service_user_v1

import (
	resource_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/resource"
	dto_service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service/dto"
)

type IUserService interface {
	CreateUser(payload *dto_service_user_v1.CreateUserDTO) error
	GetUser(payload *dto_service_user_v1.GetUserDTO) (*dto_service_user_v1.UserDTO, error)
}

type UserService struct {
	UserResource resource_user_v1.IUserResource
}

func InitUserService(userResource resource_user_v1.IUserResource) IUserService {
	return &UserService{
		UserResource: userResource,
	}
}
