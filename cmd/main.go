package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodericusifo/mini-project-echo/libs/filter"
	my_validator "github.com/rodericusifo/mini-project-echo/libs/validator"
	v1 "github.com/rodericusifo/mini-project-echo/src/core/v1"
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
	v1.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8888"))
}
