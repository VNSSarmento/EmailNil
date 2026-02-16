package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaign, err := h.CampaingService.Get()

	return campaign, 200, err
}

func (h *Handler) CampaignGetId(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")

	campaign, err := h.CampaingService.GetId(id)

	return campaign, 200, err
}
