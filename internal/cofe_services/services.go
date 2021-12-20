package cofe_services




type CofeServices struct {
	NawMenu *NMenu
	//ToDo Тут остальные модели
}


func NewCofeService(nm *NMenu) *CofeServices {
	
	return &CofeServices{
		NawMenu: nm,
	}
}


