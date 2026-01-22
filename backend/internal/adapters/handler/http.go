package handler

import (
	"encoding/json"
	"net/http"

	"collector-library/internal/core/services"
)

type HTTPHandler struct {
	svc *services.AmiiboService
}

func NewHTTPHandler(svc *services.AmiiboService) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

func (h *HTTPHandler) ListAmiibos(w http.ResponseWriter, r *http.Request) {
	amiibos, err := h.svc.ListAmiibos(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch amiibos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(amiibos)
}
