package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"time"

	// the '_' means that we will not direct use the pq package
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type serverConfig struct {
	Port        int
	Environment string
	db          struct {
		dsn string
	}
}

type application struct {
	config  serverConfig
	logger  *slog.Logger
	version string
}

func main() {
	var settings serverConfig

	flag.IntVar(&settings.Port, "port", 4002, "Server port")
	flag.StringVar(&settings.Environment, "env", "development", "Environment(development|staging|production)")
	// read in the dsn
	flag.StringVar(&settings.db.dsn, "db-dsn", "postgres://quotes:quotes2025@localhost/quotes?sslmode=disable", "PostgreSQL DSN")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// the call to openDB() sets up our connection pool
	db, err := openDB(settings)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// release the database resources before exiting
	defer db.Close()

	logger.Info("database connection pool established")

	appInstance := &application{
		config:  settings,
		logger:  logger,
		version: version,
	}

	err = appInstance.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

} // end of main()

func openDB(settings serverConfig) (*sql.DB, error) {
	// open a connection pool
	db, err := sql.Open("postgres", settings.db.dsn)
	if err != nil {
		return nil, err
	}

	// set a context to ensure DB operations don't take too long
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// let's test if the connection pool was created
	// we trying pinging it with a 5-second timeout
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	// return the connection pool (sql.DB)
	return db, nil

}
