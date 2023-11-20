package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBMysqlOption options for mysql connection
type DBMysqlOption struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	Name                 string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

func NewMysqlDatabase(option DBMysqlOption) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", option.Username, option.Password, option.Host, option.Port, option.Name, option.AdditionalParameters)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	return
}
