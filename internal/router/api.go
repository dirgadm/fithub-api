package router

import (
	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/handler"
	"github.com/dirgadm/fithub-api/internal/middleware"
	"github.com/dirgadm/fithub-api/internal/service"
	"github.com/labstack/echo/v4"
)

func ApiRouter(e *echo.Echo, opt common.Options, services *service.Services) {
	Hopt := handler.HOption{
		Common:   opt,
		Services: services,
	}

	healthCheckHandler := handler.HealthCheckHandler{
		HOption: Hopt,
	}
	authHandler := handler.AuthHandler{
		HOption: Hopt,
	}

	taskHandler := handler.TaskHandler{
		HOption: Hopt,
	}

	mw := middleware.NewMiddleware(opt)

	v1 := e.Group("v1")
	v1.GET("/health_check", healthCheckHandler.HealthCheck)

	v1.POST("/auth/login", authHandler.Login)
	v1.POST("/auth/register", authHandler.Register)

	v1.GET("/api/task", taskHandler.GetList, mw.Authorized())
	v1.POST("/api/task", taskHandler.Create, mw.Authorized())
	v1.PUT("/api/task/:id", taskHandler.Update, mw.Authorized())

}
