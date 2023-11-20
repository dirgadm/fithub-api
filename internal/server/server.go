package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/router"
	"github.com/dirgadm/fithub-api/internal/service"
	"github.com/sirupsen/logrus"
)

// IServer interface for server
type IServer interface {
	StartRestServer()
}

type server struct {
	opt      common.Options
	services *service.Services
}

// NewServer create object server
func NewServer(opt common.Options, services *service.Services) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func (s *server) StartRestServer() {
	var srv http.Server

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logrus.Infoln("[API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logrus.Errorf("[API] Fail to shutting down: %v", err)
		}

		close(done)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", s.opt.Config.App.Host, s.opt.Config.App.Port)

	srv.Handler = router.Router(s.opt, s.services)

	logrus.Infof("[API] HTTP serve at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener
		logrus.Errorf("[API] Fail to start listen and server: %v", err)
	}

	<-done
	logrus.Info("[API] Bye")
}
