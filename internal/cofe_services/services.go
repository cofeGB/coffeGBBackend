package cofe_services

type CofeServices struct {
	NawMenu  *NMenu
	Category *CategoryRepo
	//ToDo Тут остальные модели
}

func NewCofeService(menuStore NawMenuStore, categoryStore CategoryStore) *CofeServices {
	return &CofeServices{
		NawMenu:  NewNawMenu(menuStore),
		Category: NewCategories(categoryStore),
	}
}
