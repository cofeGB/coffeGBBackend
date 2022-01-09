package cofe_api

import (
	
	"net/http"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
)

type CofeHandler struct {
	service cofe_services.CofeServices
}

func (h *CofeHandler) Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}