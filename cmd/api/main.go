package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const appVersion = "1.0.0"

type serverConfig struct {
	Port        int
	Environment string
}

type application struct {
	config serverConfig
	logger *slog.Logger
}

func main() {
	var settings serverConfig

	flag.IntVar(&settings.Port, "port", 4002, "Server port")
	flag.StringVar(&settings.Environment, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	appInstance := &application{
		config: settings,
		logger: logger,
	}

	// router := http.NewServeMux()
	// router.HandleFunc("/v1/healthcheck", appInstance.healthcheckHandler)

	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", settings.Port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "address", apiServer.Addr,
		"environment", settings.Environment)
	err := apiServer.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
