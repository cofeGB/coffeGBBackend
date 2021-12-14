package starterstore

import (
	"log"

	"gorm.io/driver/sqlite"

	//"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/navmenu"

	"github.com/google/uuid"
)

func StartDB(dbname string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	_ = migrateNavMenu(db)

	appendDataNavMenu(db)

	return db, err

}

func migrateNavMenu(db *gorm.DB) error {
	if err := db.AutoMigrate(&navmenu.NavMenu{}); err != nil {
		return err
	}
	return nil

}

func appendDataNavMenu(db *gorm.DB) {
	var nm = []navmenu.NavMenu{
		{
			ID:      uuid.New(),
			Title:   "Закуски",
			Route:   "/menu",
			Icon:    "starters",
			Query:   "starters",
			Warn:    false,
			WarnMsg: "",
		},
		{
			ID:      uuid.New(),
			Title:   "Сендвичи",
			Route:   "/menu",
			Icon:    "sandwich",
			Query:   "sandwich",
			Warn:    false,
			WarnMsg: "",
		},
		{
			ID:      uuid.New(),
			Title:   "Салаты",
			Route:   "/menu",
			Icon:    "salad",
			Query:   "salad",
			Warn:    false,
			WarnMsg: "",
		},
		{
			ID:      uuid.New(),
			Title:   "Десерты",
			Route:   "/menu",
			Icon:    "desserts",
			Query:   "desserts",
			Warn:    false,
			WarnMsg: "",
		}}
	db.Create(&nm)
}
