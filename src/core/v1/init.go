package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/mini-project-echo/wire"
)

func InitRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		v1User := v1.Group("/users")
		v1UserController := wire.UserControllerV1()
		v1UserController.Mount(v1User)
	}
}
