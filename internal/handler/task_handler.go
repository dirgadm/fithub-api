package handler

import (
	"github.com/dirgadm/fithub-api/internal/dto"
	"github.com/dirgadm/fithub-api/pkg/ehttp"
	"github.com/dirgadm/fithub-api/pkg/log"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	HOption
}

func (h TaskHandler) GetList(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	// get pagination
	var page *ehttp.Paginator
	page, err = ehttp.NewPaginator(ctx)
	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	// get params
	search := ctx.GetParamString("search")

	var products []dto.TaskResponse
	var total int64
	products, total, err = h.Services.Task.GetList(ctx.Request().Context(), page.Start, page.Limit, search)
	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	ctx.DataList(products, total, page.Page, page.PerPage)

	return ctx.Serve(err)
}

func (h TaskHandler) Create(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var req dto.CreateTaskRequest

	if err = ctx.Bind(&req); err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	if err = h.Common.Validate.Struct(req); err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	ctx.ResponseData, err = h.Services.Task.Create(ctx.Request().Context(), req)
	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	return ctx.Serve(err)
}

func (h TaskHandler) Update(c echo.Context) (err error) {
	ctx := c.(*ehttp.Context)

	var id int
	id, err = ctx.GetParamUri("id")
	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	var tasks dto.TaskResponse
	tasks, err = h.Services.Task.Update(ctx.Request().Context(), id)
	if err != nil {
		h.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return ctx.Serve(err)
	}

	ctx.Data(tasks)

	return ctx.Serve(err)
}
