package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v10"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mauricioww/goauth/app/repository"
	"github.com/mauricioww/goauth/app/service"
	"github.com/mauricioww/goauth/app/transport"
)

type dbConfig struct {
	User     string `env:"MYSQL_USER" envDefault:"admin"`
	Password string `env:"MYSQL_PASSWORD" envDefault:"password"`
	Name     string `env:"MYSQL_DATABASE" envDefault:"goauth"`
	Host     string `env:"DB_HOST" envDefault:"goauth_db"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
}

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(
		logger,
		"service",
		"goauth",
		"time",
		log.DefaultTimestampUTC,
		"caller",
		log.DefaultCaller,
	)

	level.Info(logger).Log("message", "Service running...")
	defer level.Info(logger).Log("message", "Service stopped")

	dbCfg := dbConfig{}
	if err := env.Parse(&dbCfg); err != nil {
		level.Error(logger).Log("Fatal: ", err)
		os.Exit(-1)
	}

	mySQLAddress := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	level.Info(logger).Log("message", fmt.Sprintf("Connecting %v", mySQLAddress))

	mySQL, err := sql.Open("mysql", mySQLAddress)
	if err != nil {
		level.Error(logger).Log("Fatal: ", err)
		os.Exit(-1)
	}

	repository := repository.NewRepository(mySQL, logger)
	service := service.NewService(repository, logger)
	endpoints := transport.MakeHttpEndpoints(service)

	errors := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		handler := transport.NewHttpServer(endpoints)
		level.Info(logger).Log("Info:", "Server running on :8080")
		errors <- http.ListenAndServe(":8080", handler)
	}()

	level.Error(logger).Log("exit: ", <-errors)
}
