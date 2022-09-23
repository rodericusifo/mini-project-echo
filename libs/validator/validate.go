package validator

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type IRequest interface {
	CustomValidate() error
}

func ValidateRequest(ctx echo.Context, req IRequest) error {
	if err := ctx.Bind(req); err != nil {
		logrus.Infoln("bind request data error", err)
		return err
	}
	if err := ctx.Validate(req); err != nil {
		logrus.Infoln("validate request data error", err)
		return err
	}
	if err := req.CustomValidate(); err != nil {
		return err
	}
	return nil
}
