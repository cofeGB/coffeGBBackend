// Package main implements CLI to setup and run backend server on Heroku platform.
// Makes some specific actions for Heroku Proxy, DB integrations
package main

import (
	//std
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	// third party
	"github.com/kelseyhightower/envconfig"

	// my own
	"github.com/cofeGB/coffeGBBackend/internal/cofe_api"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage"
)

const (
	envPrefix             = "COFFEGB"
	herokuNginxSignalFile = "/tmp/app-initialized"
)

type ServerSettings struct {
	Listen   string `default:"127.0.0.1:8123"`
	LogLevel string `default:"INFO"`
	DBFile   string `default:"coffeDb.db"`
}

func setUp() (srv *ServerSettings) {
	flag.Parse()

	srv = &ServerSettings{}
	// always try to read env, maybe use defaults
	if err := envconfig.Process(envPrefix, srv); err != nil {
		log.Fatalln(err)
	}

	return srv
}

func main() {
	// register signal handlers
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// setup app
	srvSetting := setUp()

	// init storages and services

	storage, err := cofe_storage.NewCofeStorage(srvSetting.DBFile)
	if err != nil {
		log.Fatalf("cannot initialize storage: %s", err.Error())
	}
	defer storage.Close()

	cofeService := cofe_services.NewCofeService(storage)

	// start api server
	server := cofe_api.NewCofeAPIServer(srvSetting.Listen, srvSetting.LogLevel, cofeService)

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