package handler

import (
	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/service"
)

type HOption struct {
	Common   common.Options
	Services *service.Services
}
