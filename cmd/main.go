package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodericusifo/mini-project-echo/libs/filter"
	"github.com/rodericusifo/mini-project-echo/src/core/v1"
	my_validator "github.com/rodericusifo/mini-project-echo/libs/validator"
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
	v1.NewRoutes(e)

	e.Logger.Fatal(e.Start(":8888"))
}
