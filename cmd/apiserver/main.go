// Package main implements CLI to setup and run backend server on Heroku platform.
// Makes some specific actions for Heroku Proxy, DB integrations
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_api"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage"
)

const (
	envPrefix             = "COFFEGB"
	envDbPrefix           = "DATABASE"
	herokuNginxSignalFile = "/tmp/app-initialized"
)

type DBSettings struct {
	URL           string        `default:"host=localhost port=5432 user=postgres password=postgres dbname=cofeGB sslmode=disable"`
	QueryTimeout  time.Duration `default:"30s"`
	MigrationsDir string        `default:"migrations"`
}

type ServerSettings struct {
	Listen   string `default:"127.0.0.1:8123"`
	LogLevel string `default:"INFO"`
}

func setUp() (srv *ServerSettings, db *DBSettings) {
	srv = &ServerSettings{}
	db = &DBSettings{}
	flag.Usage = func() {
		fmt.Print("-- App server config --\n\n")
		_ = envconfig.Usage(envPrefix, srv)
		fmt.Print("\n-- Database config --\n\n")
		_ = envconfig.Usage(envDbPrefix, db)
	}
	flag.Parse()

	// always try to read env, maybe use defaults
	if err := envconfig.Process(envPrefix, srv); err != nil {
		log.Fatalln(err)
	}
	if err := envconfig.Process(envDbPrefix, db); err != nil {
		log.Fatalln(err)
	}

	return srv, db
}

func main() {
	// register signal handlers
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// setup app
	srvSetting, dbSettings := setUp()

	// init store db
	db, err := cofe_storage.NewCofeStore(&cofe_storage.Settings{
		DSN:           dbSettings.URL,
		MigrationsDir: dbSettings.MigrationsDir,
	})
	if err != nil {
		log.Fatalf("cannot initialize storage: %s", err.Error())
	}
	defer db.PG.Close()

	// init services
	cofeServices := cofe_services.NewCofeService(
		cofe_storage.NewNawMenuStore(db),
		cofe_storage.NewCategoryStore(db),
		cofe_storage.NewDishStore(db),
	)

	// start api server
	server := cofe_api.NewCofeAPIServer(srvSetting.Listen, srvSetting.LogLevel, *cofeServices)

	go func() {
		// usually server works behind proxy,
		// so just run plain http server
		err := server.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	// touch file to signal that nginx can start
	_, err = os.Stat(herokuNginxSignalFile)
	if os.IsNotExist(err) {
		f, err := os.Create(herokuNginxSignalFile)
		if err != nil {
			log.Fatalf("cannot create %s file for nginx integration, err: %s", herokuNginxSignalFile, err.Error())
		}
		f.Close()
	} else {
		err = os.Chtimes(herokuNginxSignalFile, time.Now().Local(), time.Now().Local())
		if err != nil {
			log.Fatalf("cannot touch %s file for nginx integration, err: %s", herokuNginxSignalFile, err.Error())
		}
	}

	// wait signal to gracefully shutdown
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}
