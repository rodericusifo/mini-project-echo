package dto_service_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/libs/dto"
	request_controller_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/controller/request"
)

type GetUserDTO struct {
	dto.DTO[*request_controller_user_v1.GetUserRequestParam, any, any]
}
