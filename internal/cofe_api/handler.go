package cofe_api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/services"
)	

type CofeHandler struct {
	service services.CofeServices
}

func (h *CofeHandler) Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}

//http://localhost:3000/api/navMenu/
/**
 * @typedef {Object} NavMenuItem - элемент навигационного меню приложения
 * @property {string} title - название пункта навигационного меню, как его видит пользователь
 * @property {string} path - полный путь роутера '/menu/sandwiches'|'/menu/salads'...
 * @property {number} itemOrder - порядковый номер данного элемента в меню, целые числа 0,1,2... 
 */

 type NavMenuItem struct {
	Title string `json:"title"`
	Path string `json:"path"`
	ItemOrder int32 `json:"itemOrder"`
	

 }
func (h *CofeHandler)GetListNawMenu(w http.ResponseWriter, r *http.Request) {
	//_, _ = w.Write([]byte("ListNawMenu"))

	nm, err :=h.service.NawMenu.GetListNawMenu(r.Context())  
	if err != nil {
		log.Printf("")
	}
	 nmi :=NavMenuItem{}
	listNmi :=[]NavMenuItem{}
	for _, mi := range nm {
		nmi.Title = mi.Title
		nmi.Path = mi.Path
		nmi.ItemOrder = mi.ItemOrder
		listNmi = append(listNmi, nmi)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(listNmi)


}