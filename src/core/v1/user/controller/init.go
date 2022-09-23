package controller_user_v1

import (
	"github.com/labstack/echo/v4"
	service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service"
)

type UserController struct {
	UserService service_user_v1.IUserService
}

func InitUserController(userService service_user_v1.IUserService) *UserController {
	return &UserController{UserService: userService}
}

func (userController *UserController) Mount(group *echo.Group) {
	group.POST("/create", userController.CreateUser)
	group.GET("/:id",userController.GetUser )
}
