package main

import (
	"flag"
	"log/slog"
	"os"
)

// configuration defines the application's configuration settings.
type configuration struct {
	port int
	env  string
}

// application holds the application's dependencies.
type application struct {
	config configuration
	logger *slog.Logger
}

func main() {
	cfg := loadConfig()
	logger := setupLogger(cfg.env)

	app := &application{
		config: cfg,
		logger: logger,
	}

	if err := app.serve(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

// loadConfig reads configuration from command-line flags.
func loadConfig() configuration {
	var cfg configuration

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	return cfg
}

// setupLogger configures the application logger based on the environment.
func setupLogger(env string) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}
