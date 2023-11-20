package common

import (
	"database/sql"

	"github.com/dirgadm/fithub-api/config"
	"github.com/dirgadm/fithub-api/pkg/log"
	"github.com/go-playground/validator/v10"
)

// Options common option for all object that needed
type Options struct {
	Config   config.Config
	Database Database
	Logger   *log.Logger
	Validate *validator.Validate
}

type Database struct {
	Read  *sql.DB
	Write *sql.DB
}
