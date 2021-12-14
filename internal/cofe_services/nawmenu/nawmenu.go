package nawmenu

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type NawMenu struct {
	ID uuid.UUID
	Title string
	Route string
	Icon string
	Query string
	Warn bool
	WarnMsg string 
}

type NawMenuStore interface {
	ReadNawMenu(ctx context.Context) ([]NawMenu, error)
}


type NMenu struct {
	 nawMenuStore NawMenuStore
}
func NewMenu(nawMenuStore NawMenuStore ) *NMenu {
	return &NMenu{
		nawMenuStore: nawMenuStore,
	}
}

func(ms *NMenu) ReadMenu(ctx context.Context) ([]NawMenu, error) {
 
	return nil,nil
}