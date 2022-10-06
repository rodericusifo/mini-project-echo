package service_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/libs/util"
	request_controller_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/controller/request"
	dto_service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service/dto"
	"github.com/rodericusifo/mini-project-echo/src/model"
)

func (s *UserService) CreateUser(payload *dto_service_user_v1.CreateUserDTO) error {
	userModel, err := util.ConvertToStruct[model.User](struct {
		*request_controller_user_v1.CreateUserRequestBody
	}{payload.Body})
	if err != nil {
		return err
	}

	err = s.UserRepository.CreateUser(&userModel)
	if err != nil {
		return err
	}

	return nil
}
