package config

import (
	"fmt"
	"time"

	"github.com/dirgadm/fithub-api/pkg/env"
	"github.com/go-playground/validator"
)

type (
	// config
	Config struct {
		// app
		App struct {
			Host  string `validate:"required"`
			Port  int    `validate:"required"`
			Name  string `validate:"required"`
			Debug bool   `validate:"required"`
			Env   string `validate:"required"`
		}

		// jwt
		Jwt struct {
			Key string `validate:"required"`
		}

		// database
		Database struct {
			Enabled         bool
			Connection      string        `validate:"required"`
			Name            string        `validate:"required"`
			MaxOpenConns    int           `validate:"required"`
			MaxIdleConns    int           `validate:"required"`
			ConnMaxLifetime time.Duration `validate:"required"`
			// database.write
			Write struct {
				Host     string `validate:"required"`
				Port     int    `validate:"required"`
				Username string `validate:"required"`
				Password string `validate:"required"`
			}
			// database.read
			Read struct {
				Host     string `validate:"required"`
				Port     int    `validate:"required"`
				Username string `validate:"required"`
				Password string `validate:"required"`
			}
		}
	}
)

// NewConfig returns app config.
func NewConfig() *Config {
	cfg := &Config{}
	env, err := env.Env("env")
	if err != nil {
		fmt.Printf("Failed to using config file, error find %s | %v ", env.ConfigFileUsed(), err)
		return cfg
	}

	// validator
	validate := validator.New()

	// app
	cfg.App.Host = env.GetString("app.host")
	cfg.App.Port = env.GetInt("app.port")
	cfg.App.Name = env.GetString("app.name")
	cfg.App.Debug = env.GetBool("app.debug")
	cfg.App.Env = env.GetString("app.env")

	// jwt
	cfg.Jwt.Key = env.GetString("jwt.key")

	// database
	cfg.Database.Enabled = env.GetBool("database.enabled")
	cfg.Database.Connection = env.GetString("database.connection")
	cfg.Database.Name = env.GetString("database.name")
	cfg.Database.MaxOpenConns = env.GetInt("database.max_open_conns")
	cfg.Database.MaxIdleConns = env.GetInt("database.max_idle_conns")
	cfg.Database.ConnMaxLifetime = env.GetDuration("database.conn_lifetime_max")
	// database.write
	cfg.Database.Write.Host = env.GetString("database.write.host")
	cfg.Database.Write.Port = env.GetInt("database.write.port")
	cfg.Database.Write.Username = env.GetString("database.write.username")
	cfg.Database.Write.Password = env.GetString("database.write.password")
	// database.read
	cfg.Database.Read.Host = env.GetString("database.read.host")
	cfg.Database.Read.Port = env.GetInt("database.read.port")
	cfg.Database.Read.Username = env.GetString("database.read.username")
	cfg.Database.Read.Password = env.GetString("database.read.password")
	// validate config
	if cfg.Database.Enabled {
		err = validate.Struct(cfg.Database)
		if err != nil {
			fmt.Printf("Failed to find specific config file, error find %s | %v ", env.ConfigFileUsed(), err)
			return cfg
		}
	}

	return cfg
}
