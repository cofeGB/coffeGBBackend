package cofe_storage

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type CofeDB struct {
	PG *sql.DB
}

type Settings struct {
	DSN           string
	MigrationsDir string
}

var dbConn *CofeDB = nil

func NewCofeStore(settings *Settings) (*CofeDB, error) {
	if dbConn != nil {
		return dbConn, nil
	}

	db, err := sql.Open("postgres", settings.DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:"+settings.MigrationsDir,
		"postgres", driver)
	m.Up()
	if err != nil {
		log.Fatalf("cannot run migration: %s", err.Error())
	}

	return &CofeDB{PG: db}, nil
}

func (cs *CofeDB) Close() {

	cs.PG.Close()
}
