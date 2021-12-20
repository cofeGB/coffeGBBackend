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



func NewCofeStore(dsn string) (*CofeDB, error) {
	db, err := sql.Open("postgres", dsn)
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
        "file:migrations",
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
