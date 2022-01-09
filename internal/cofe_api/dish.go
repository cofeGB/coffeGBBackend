package cofe_api

import (
	"encoding/json"
	"fmt"

	//	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	//"github.com/docker/docker/api/server/router"
	//"github.com/google/uuid"
)

func (h *CofeHandler) GetListDish(w http.ResponseWriter, r *http.Request) {
	d, err := h.service.Dish.GetListDIsh(r.Context())
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_ = json.NewEncoder(w).Encode(d)

}

func (h *CofeHandler) Menu(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			if len(r.URL.Query()) > 1 {
				fmt.Fprintf(w, "number of parameters in the request > 1 ")
				return
			}
			if cat := r.URL.Query().Get("category"); cat != "" {
				d, err := h.service.Dish.GetDishByCategory(r.Context(), cat)
				if err != nil {
					log.Println(err)
					return
				}
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				_ = json.NewEncoder(w).Encode(d)
				return
			}

			if dish := r.URL.Query().Get("dish"); dish != "" {
				dUUID, err := uuid.Parse(dish)
				if err != nil {
					log.Println(err)
					return
				}
				d, err := h.service.Dish.GetDIshByID(r.Context(), dUUID)
				if err != nil {
					log.Println(err)
					return
				}
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				_ = json.NewEncoder(w).Encode(d)
				return
			}

		}
	case "POST":
		log.Println("POST")
	}

}
