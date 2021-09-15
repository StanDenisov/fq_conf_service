package router

import (
	"github.com/StanDenisov/fq_conf_service/controller"
	"github.com/labstack/echo"
)

//Router - default logic router
func Router(e *echo.Echo) {
	e.POST("/", controller.SendConf)
}
