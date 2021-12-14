package navmenu

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type NavMenu struct {
	ID      uuid.UUID
	Title   string
	Route   string
	Icon    string
	Query   string
	Warn    bool
	WarnMsg string
}

type NavMenuStore interface {
	ReadNavMenu(ctx context.Context) ([]NavMenu, error)
}

type NMenu struct {
	navMenuStore NavMenuStore
}

func NewMenu(navMenuStore NavMenuStore) *NMenu {
	return &NMenu{
		navMenuStore: navMenuStore,
	}
}

func (ms *NMenu) ReadMenu(ctx context.Context) ([]NavMenu, error) {

	return nil, nil
}
