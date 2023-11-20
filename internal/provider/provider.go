package provider

import (
	"database/sql"

	"github.com/dirgadm/fithub-api/config"
	"github.com/dirgadm/fithub-api/pkg/driver"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
)

const (
	DBDialectMysql = "mysql"
)

// AppContext the app context struct
type AppProvider struct {
	config config.Config
}

// NewAppContext initiate appcontext object
func NewAppProvider(config config.Config) *AppProvider {
	return &AppProvider{
		config: config,
	}
}

func (a *AppProvider) GetDBInstanceRead() (sqlDB *sql.DB, err error) {
	dbOption := driver.DBMysqlOption{
		Host:                 a.config.Database.Read.Host,
		Port:                 a.config.Database.Read.Port,
		Username:             a.config.Database.Read.Username,
		Password:             a.config.Database.Read.Password,
		Name:                 a.config.Database.Name,
		AdditionalParameters: "parseTime=true",
	}

	sqlDB, err = driver.NewMysqlDatabase(dbOption)
	return
}

func (a *AppProvider) GetDBInstanceWrite() (sqlDB *sql.DB, err error) {
	dbOption := driver.DBMysqlOption{
		Host:                 a.config.Database.Write.Host,
		Port:                 a.config.Database.Write.Port,
		Username:             a.config.Database.Write.Username,
		Password:             a.config.Database.Write.Password,
		Name:                 a.config.Database.Name,
		AdditionalParameters: "parseTime=true",
	}

	sqlDB, err = driver.NewMysqlDatabase(dbOption)
	return
}

func (a *AppProvider) GetDBDriver(sqlDB *sql.DB) (dbDriver database.Driver, err error) {
	dbDriver, err = mysql.WithInstance(sqlDB, &mysql.Config{})
	return
}
