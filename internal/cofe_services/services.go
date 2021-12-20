package cofe_services

type CofeServices struct {
	NawMenu *NMenu
	//ToDo Тут остальные модели
}

func NewCofeService(menuStore NawMenuStore) *CofeServices {
	nm := NewNawMenu(menuStore)

	return &CofeServices{
		NawMenu: nm,
	}
}
