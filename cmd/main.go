package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/mini-project-echo/libs/filter"
	my_validator "github.com/rodericusifo/mini-project-echo/libs/validator"
	"github.com/rodericusifo/mini-project-echo/wire"
)

func main() {
	e := echo.New()
	e.Validator = &my_validator.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = filter.CustomHTTPErrorHandler

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	// Version 1
	v1 := e.Group("/v1")

	v1User := v1.Group("/users")
	v1UserController := wire.UserControllerV1()
	v1UserController.Mount(v1User)

	e.Logger.Fatal(e.Start(":8888"))
}
