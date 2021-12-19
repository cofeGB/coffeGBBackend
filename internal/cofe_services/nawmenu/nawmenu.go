package nawmenu

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type NawMenu struct {
	ID uuid.UUID
	Title string
	Path string
	ItemOrder int32
	Is_run    bool
}





type NawMenuStore interface {
	GetListNawMenu(ctx context.Context) ([]NawMenu, error)
}


type NMenu struct {
	 Store NawMenuStore
}
func NewNawMenu(nawMenuStore NawMenuStore ) *NMenu {
	return &NMenu{
		Store: nawMenuStore,
	}
}

func(ms *NMenu) GetListNawMenu(ctx context.Context) ([]NawMenu, error) {
    nm, err:= ms.Store.GetListNawMenu(ctx)

	if err != nil {
		return nil, err
	}
	return nm,nil
}