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

func (c *UserController) CreateUser(ctx echo.Context) error {
	reqBody := new(request_controller_user_v1.CreateUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	if err := c.UserService.CreateUser(&dto_service_user_v1.CreateUserDTO{
		DTO: dto.DTO[any, any, *request_controller_user_v1.CreateUserRequestBody]{
			Param: nil,
			Query: nil,
			Body:  reqBody,
		},
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.ResponseSuccess[any]("success create user", nil))
}
