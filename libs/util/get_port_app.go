package util

import (
	"strconv"

	"github.com/rodericusifo/mini-project-echo/libs/config"
)

func GetPortApp() string {
	return strconv.Itoa(config.AppConfig.Server.Port)
}
