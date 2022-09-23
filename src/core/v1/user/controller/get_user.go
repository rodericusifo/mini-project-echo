package controller_user_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/mini-project-echo/libs/dto"
	"github.com/rodericusifo/mini-project-echo/libs/response"
	"github.com/rodericusifo/mini-project-echo/libs/validator"
	request_controller_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/controller/request"
	dto_service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service/dto"
)

func (c *UserController) GetUser(ctx echo.Context) error {
	reqParam := new(request_controller_user_v1.GetUserRequestParam)
	if err := validator.ValidateRequest(ctx, reqParam); err != nil {
		return err
	}

	userDto, err := c.UserService.GetUser(&dto_service_user_v1.GetUserDTO{
		DTO: dto.DTO[*request_controller_user_v1.GetUserRequestParam, any, any]{
			Param: reqParam,
			Query: nil,
			Body:  nil,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("Get user success", userDto))
}
