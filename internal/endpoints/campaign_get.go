package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaign, err := h.CampaingService.Repository.Get()

	return campaign, 200, err
}
