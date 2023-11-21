package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/dirgadm/fithub-api/config"
	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/provider"
	"github.com/dirgadm/fithub-api/internal/repository"
	"github.com/dirgadm/fithub-api/internal/server"
	"github.com/dirgadm/fithub-api/internal/service"
	"github.com/dirgadm/fithub-api/pkg/log"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.NewConfig()
	prv := provider.NewAppProvider(*cfg)

	var err error

	lgr := logrus.New()
	if cfg.App.Debug {
		lgr.SetFormatter(log.NewFormater(true, cfg.App.Name))
	} else {
		lgr.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}
	lgr.SetReportCaller(true)
	lgr.SetOutput(os.Stdout)

	logger := log.NewLoggerWithClient(cfg.App.Name, lgr)

	var dbRead *sql.DB
	dbRead, err = prv.GetDBInstanceRead()
	if err != nil {
		logrus.Errorf("Failed to start, error connect to database | %v", err)
		return
	}

	err = dbRead.Ping()
	if err != nil {
		logrus.Fatalf("Failed to get database instance | %v", err)
		return
	}
	defer dbRead.Close()

	var dbWrite *sql.DB
	dbWrite, err = prv.GetDBInstanceWrite()
	if err != nil {
		logrus.Errorf("Failed to start, error connect to database | %v", err)
		return
	}

	err = dbWrite.Ping()
	if err != nil {
		logrus.Fatalf("Failed to get database instance | %v", err)
		return
	}
	defer dbWrite.Close()

	validate := validator.New()

	opt := common.Options{
		Config: *cfg,
		Database: common.Database{
			Read:  dbRead,
			Write: dbWrite,
		},
		Logger:   logger,
		Validate: validate,
	}

	repo := wiringRepository(repository.ROption{
		Common: opt,
	})

	svc := wiringService(service.SOption{
		Common:     opt,
		Repository: repo,
	})

	// run app
	svr := server.NewServer(opt, svc)
	svr.StartRestServer()
}

func wiringRepository(repoOption repository.ROption) *repository.Repository {
	user := repository.NewUser(repoOption)
	task := repository.NewTask(repoOption)

	repo := repository.Repository{
		User: user,
		Task: task,
	}

	return &repo
}

func wiringService(serviceOption service.SOption) *service.Services {
	auth := service.NewAuthService(serviceOption)
	task := service.NewTaskService(serviceOption)

	svc := service.Services{
		Auth: auth,
		Task: task,
	}

	return &svc
}
