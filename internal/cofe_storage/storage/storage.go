package storage

import (
	"log"

	"gorm.io/gorm"
	
)

type CofeDB struct {
	DB *gorm.DB
}

func NewCofeStore(db *gorm.DB) (*CofeDB, error) {
	return &CofeDB{DB: db}, nil
}

func (c *CofeDB) Close() {
	sqlDB, err := c.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
