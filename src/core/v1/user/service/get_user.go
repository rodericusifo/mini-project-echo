package service_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/libs/util"
	request_controller_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/controller/request"
	dto_service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service/dto"
	"github.com/rodericusifo/mini-project-echo/src/model"
)

func (s *UserService) GetUser(payload *dto_service_user_v1.GetUserDTO) (*dto_service_user_v1.UserDTO, error) {
	userModel, err := util.ConvertToStruct[model.User](struct {
		*request_controller_user_v1.GetUserRequestParam
	}{payload.Param})
	if err != nil {
		return nil, err
	}

	userModelRes, err := s.UserResource.GetUser(nil, &userModel)
	if err != nil {
		return nil, err
	}

	userDto, err := util.ConvertToStruct[*dto_service_user_v1.UserDTO](userModelRes)
	if err != nil {
		return nil, err
	}

	return userDto, nil
}
