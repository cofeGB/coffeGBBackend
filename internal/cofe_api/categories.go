package cofe_api

import (
	"encoding/json"
	"log"
	"net/http"
)

//http://localhost:3000/categories/
/**
 * @typedef {Object} NavMenuItem - элемент навигационного меню приложения
 * @property {string} title - название пункта навигационного меню, как его видит пользователь
 * @property {string} query - ??
 * @property {string} icon  - ??
 * @property {number} itemOrder - порядковый номер данного элемента, целые числа 0,1,2...
 */

func (h *CofeHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.Category.GetCategories(r.Context())
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(items)

}
