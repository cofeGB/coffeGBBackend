package services

import "github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"


type CofeServices struct {
	nawMenu *nawmenu.NMenu
	//ToDo Тут остальные модели
}


func NewCofeService(nm *nawmenu.NMenu) *CofeServices {
	
	return &CofeServices{
		nawMenu: nm,
	}
}


