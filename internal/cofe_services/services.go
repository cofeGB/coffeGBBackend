package cofe_services

type CofeServices struct {
	NawMenu  *NMenu
	Category *CategoryRepo
	Dish *Dishs
	//ToDo Тут остальные модели
}

func NewCofeService(menuStore NawMenuStore, categoryStore CategoryStore,dishStore DishStore ) *CofeServices {
	return &CofeServices{
		NawMenu:  NewNawMenu(menuStore),
		Category: NewCategoryRepo(categoryStore),
		Dish: NewDishs(dishStore),
	}
}
