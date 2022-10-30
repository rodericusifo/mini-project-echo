//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	controller_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/controller"
	resource_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/resource"
	service_user_v1 "github.com/rodericusifo/mini-project-echo/src/core/v1/user/service"
)

func UserControllerV1() *controller_user_v1.UserController {
	wire.Build(
		controller_user_v1.InitUserController,
		service_user_v1.InitUserService,
		resource_user_v1.InitUserResource,
	)
	return &controller_user_v1.UserController{}
}
