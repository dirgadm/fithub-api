package service

import (
	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/repository"
)

// SOption anything any service object needed
type SOption struct {
	Common     common.Options
	Repository *repository.Repository
}

// Services all service object injected here
type Services struct {
	Auth IAuthService
	Task ITaskService
}
