package navmenustore

import (
	"context"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/navmenu"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ navmenu.NavMenuStore = &NavMenu{}

type GormNavMenu struct {
	gorm.Model
	ID      uuid.UUID
	Title   string
	Router  string
	Icon    string
	Query   string
	Warn    bool
	WarnMsg string
}

type NavMenu struct {
	DB *gorm.DB
}

func NewNavMenu(db *gorm.DB) *NavMenu {
	return &NavMenu{
		DB: db,
	}
}

func (nm *NavMenu) ReadNavMenu(ctx context.Context) ([]navmenu.NavMenu, error) {
	var menuList []GormNavMenu
	result := nm.DB.Find(&menuList)
	if result.Error != nil {
		return nil, result.Error
	}
	var m navmenu.NavMenu
	var mm []navmenu.NavMenu
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
