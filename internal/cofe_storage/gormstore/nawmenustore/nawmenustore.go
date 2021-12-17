package nawmenustore

import (
	"context"
	"time"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage/gormstore/starterstore"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ nawmenu.NawMenuStore = &NawMenuStor{}

type NawMenu struct {
	ID      	uuid.UUID
	CreatedAt   time.Time	`db:"created_at"`
	UpdatedAt   time.Time	`db:"updated_at"`
	DeletedAt   time.Time	`db:"deleted_at"`
	Title 		string		`db:"title"`
	Path		string		`db:"path"`
	ItemOrder	int32		`db:"item_order"`
}

type NawMenuStor struct {
	storeGorm starterstore.StoreGorm 
}

func NewNawMenu(db starterstore.StoreGorm) *NawMenuStor {
	return &NawMenuStor{
		storeGorm: db,
	}
}

func (nm *NawMenuStor) ReadNawMenu(ctx context.Context) ([]nawmenu.NawMenu, error) {
    var NawMenuList []NawMenu
	result := nm.storeGorm.DB.Order("item_order").Find(&NawMenuList)
	if result.Error != nil {
		return nil, result.Error
	}
	var m nawmenu.NawMenu
	var mm []nawmenu.NawMenu
	for _, mg := range NawMenuList{
		m.ID = mg.ID
		m.Title = mg.Title
		m.Path = mg.Path
		m.ItemOrder = mg.ItemOrder
		mm = append(mm, m)
	}
	return mm, nil
}
