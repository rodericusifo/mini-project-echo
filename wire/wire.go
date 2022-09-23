//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/rodericusifo/mini-project-echo/libs/database"
	controller_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/controller"
	repository_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/repository"
	service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service"
)

func UserControllerV1() *controller_user_v1.UserController {
	wire.Build(
		controller_user_v1.InitUserController,
		service_user_v1.InitUserService,
		repository_user_v1.InitUserRepository,
		database.InitDatabase,
	)
	return &controller_user_v1.UserController{}
}
