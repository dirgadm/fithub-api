package router

import (
	"net/http"
	"os"
	"time"

	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/service"
	"github.com/dirgadm/fithub-api/pkg/ehttp"
	"github.com/dirgadm/fithub-api/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func Router(opt common.Options, services *service.Services) *echo.Echo {
	e := echo.New()

	// setup cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header"},
	}))

	// setup echo for request id
	e.Use(middleware.RequestID())

	// setup echo for secure
	e.Use(middleware.Secure())

	// setup echo for gzip compres
	e.Use(middleware.Gzip())

	// setup echo for recover
	if !opt.Config.App.Debug {
		e.Use(middleware.Recover())
	}

	// setup echo for real ip
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()

	// setup custom context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ehttp.Context{
				Context:        c,
				ResponseFormat: ehttp.NewResponse(),
				ResponseData:   nil,
			}
			return next(cc)
		}
	})

	// setup echo for logger
	lgr := opt.Logger.Logger()
	if opt.Config.App.Debug {
		lgr.SetFormatter(log.NewFormater(true, opt.Config.App.Name))
	} else {
		lgr.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}
	lgr.SetReportCaller(true)
	lgr.SetOutput(os.Stdout)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogURIPath:       true,
		LogRoutePath:     true,
		LogStatus:        true,
		LogLatency:       true,
		LogProtocol:      true,
		LogMethod:        true,
		LogRequestID:     true,
		LogError:         true,
		LogContentLength: true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogValuesFunc: func(ctx echo.Context, values middleware.RequestLoggerValues) (err error) {
			fields := logrus.Fields{
				"service":    opt.Config.App.Name,
				"request_id": values.RequestID,
				"uri":        values.URI,
				"uri_path":   values.URIPath,
				"route_path": values.RoutePath,
				"status":     values.Status,
				"host":       values.Host,
				"protocol":   values.Protocol,
				"method":     values.Method,
				"remote_ip":  values.RemoteIP,
				"start_time": values.StartTime,
				"end_time":   time.Now(),
				"latency":    values.Latency,
			}

			if values.Error != nil {
				fields["error"] = values.Error
			}

			if values.Status >= 400 || values.Error != nil {
				lgr.WithFields(fields).Error("HTTP " + http.StatusText(values.Status))
			} else {
				lgr.WithFields(fields).Info("HTTP " + http.StatusText(values.Status))
			}

			return
		},
	}))

	// setup binder
	e.Binder = &ehttp.Binder{}

	// setup error handler
	e.HTTPErrorHandler = ehttp.HTTPErrorHandler

	ApiRouter(e, opt, services)
	return e
}
