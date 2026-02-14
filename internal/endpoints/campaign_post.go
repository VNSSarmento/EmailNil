package endpoints

import (
	"emailNil/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaing(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	//sempre pego o meu contrato e atribuo a request
	var request contract.NewCampaign
	//faço a comparação do corpo da requisição e do contrato
	err := render.DecodeJSON(r.Body, &request)

	if err != nil {
		return nil, 401, err
	}

	//se nao tiver erro crio o novo objeto
	id, err := h.CampaingService.Create(request)

	return map[string]string{"id": id}, 201, nil
}
