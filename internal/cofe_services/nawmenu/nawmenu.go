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
}



/**
 * @typedef {Object} NavMenuItem - элемент навигационного меню приложения
 * @property {string} title - название пункта навигационного меню, как его видит пользователь
 * @property {string} path - полный путь роутера '/menu/sandwiches'|'/menu/salads'...
 * @property {number} itemOrder - порядковый номер данного элемента в меню, целые числа 0,1,2... 
 */

type NawMenuStore interface {
	GetListNawMenu(ctx context.Context) ([]NawMenu, error)
}


type NMenu struct {
	 nawMenuStore NawMenuStore
}
func NewNawMenu(nawMenuStore NawMenuStore ) *NMenu {
	return &NMenu{
		nawMenuStore: nawMenuStore,
	}
}

func(ms *NMenu) GetListNawMenu(ctx context.Context) ([]NawMenu, error) {
     
	return nil,nil
}