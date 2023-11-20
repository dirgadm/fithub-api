package handler

import (
	"net/http"

	"github.com/dirgadm/fithub-api/pkg/ehttp"
	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
	HOption
}

func (h HealthCheckHandler) HealthCheck(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	ctx.Message(http.StatusOK, "success", h.Common.Config.App.Name)
	return
}
