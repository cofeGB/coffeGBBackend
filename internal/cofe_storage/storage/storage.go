package storage

import (
	"database/sql"
	//"log"

	//"gorm.io/gorm"
)

type CofeDB struct {
	PG *sql.DB
}

// func NewCofeStore(db *gorm.DB) (*CofeDB, error) {
// 	return &CofeDB{DB: db}, nil
// }

func NewCofeStore(db *sql.DB) (*CofeDB, error) {
	return &CofeDB{PG: db}, nil
}


func (c *CofeDB) Close() {
	
	c.PG.Close()
}
