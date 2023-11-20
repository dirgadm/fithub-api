package handler

import (
	"github.com/dirgadm/fithub-api/internal/dto"
	"github.com/dirgadm/fithub-api/pkg/ehttp"
	"github.com/dirgadm/fithub-api/pkg/log"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	HOption
}

func (h AuthHandler) Login(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var req dto.LoginRequest

	if err = ctx.Bind(&req); err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	if err = h.Common.Validate.Struct(req); err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	ctx.ResponseData, err = h.Services.Auth.Login(ctx.Request().Context(), req)

	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}

func (h AuthHandler) Register(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var req dto.RegisterRequest

	if err = ctx.Bind(&req); err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	if err = h.Common.Validate.Struct(req); err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	ctx.ResponseData, err = h.Services.Auth.Register(ctx.Request().Context(), req)
	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}
