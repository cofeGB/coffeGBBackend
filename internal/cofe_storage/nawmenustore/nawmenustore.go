package nawmenustore

import (
	"context"
	"log"
	"time"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage/storage"


	"github.com/google/uuid"
)

var _ nawmenu.NawMenuStore = &NawMenuStor{}

type NawMenu struct {
	ID        uuid.UUID
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
	Title     string    `db:"title"`
	Path      string    `db:"path"`
	ItemOrder int32     `db:"item_order"`
}

type NawMenuStor struct {
	store *storage.CofeDB
}

func NewNawMenuStore(db *storage.CofeDB) *NawMenuStor {

	st := &NawMenuStor{
		store: db,
	}
	err := migrateNawMenu(st)
	if err != nil {
		log.Fatalf("cannot initialize migration NawMenu struct: %s", err.Error())
	}
	return st
}

func (nm *NawMenuStor) GetListNawMenu(ctx context.Context) ([]nawmenu.NawMenu, error) {
	var NawMenuList []NawMenu
	result := nm.store.DB.Order("item_order").Find(&NawMenuList)
	if result.Error != nil {
		return nil, result.Error
	}
	var m nawmenu.NawMenu
	var mm []nawmenu.NawMenu
	for _, mg := range NawMenuList {
		m.ID = mg.ID
		m.Title = mg.Title
		m.Path = mg.Path
		m.ItemOrder = mg.ItemOrder
		mm = append(mm, m)
	}
	return mm, nil
}

func migrateNawMenu(nm *NawMenuStor) error {
	if err := nm.store.DB.AutoMigrate(&NawMenu{}); err != nil {
		return err
	}
	return nil

}
