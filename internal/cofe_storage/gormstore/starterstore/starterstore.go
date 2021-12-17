package starterstore

import (
	"log"

	"gorm.io/driver/sqlite"

	//"github.com/mattn/go-sqlite3"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage/gormstore/nawmenustore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StoreGorm struct {
	DB *gorm.DB
	
}



func StartDB(dbname string, cfg gorm.Config) (*StoreGorm, error) {
	db, err := gorm.Open(postgres.new(dbname), &cfg)
	if err != nil {
		log.Fatal(err)
	}

	_ = migrateNawMenu(db)

	return &StoreGorm{
		DB: db,
	}, err

}

func migrateNawMenu(db *gorm.DB) error {
	if err := db.AutoMigrate(&nawmenustore.NawMenu {}); err != nil {
		return err
	}
	return nil

}

// func StartStore() *StoreGorm {
// s:=
// }
