package cofe_api

import (
	"encoding/json"
	"log"
	"net/http"

	//"github.com/google/uuid"
)










func (h *CofeHandler) GetListDish(w http.ResponseWriter, r *http.Request) {
	d, err:= h.service.Dish.GetListDIsh(r.Context())
	if err != nil {
		log.Println(err)
		return
	}
	
	

	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(d)

}