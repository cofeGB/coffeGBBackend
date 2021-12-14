package nawmenustore

import (
	"context"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ nawmenu.NawMenuStore = &NawMenu{}

type GormNawMenu struct {
	gorm.Model
	ID      uuid.UUID
	Title   string
	Router string
	Icon    string
	Query   string
	Warn    bool
	WarnMsg string
}

type NawMenu struct {
	DB *gorm.DB
}

func NewNawMenu(db *gorm.DB) *NawMenu {
	return &NawMenu{
		DB: db,
	}
}

func (nm *NawMenu) ReadNawMenu(ctx context.Context) ([]nawmenu.NawMenu, error) {
	var menuList []GormNawMenu
	result := nm.DB.Find(&menuList)
	if result.Error != nil {
		return nil, result.Error
	}
	var m nawmenu.NawMenu
	var mm []nawmenu.NawMenu
	for _, mg := range menuList {
		m.ID = mg.ID
		m.Title = mg.Title
		m.Icon = mg.Icon
		m.Query = mg.Query
		m.Warn = mg.Warn
		m.WarnMsg = mg.WarnMsg

		mm = append(mm, m)
	}
	return mm, nil
}
