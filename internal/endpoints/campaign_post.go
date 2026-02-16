package endpoints

import (
	"emailNil/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCampaing(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	err := render.DecodeJSON(r.Body, &request)

	if err != nil {
		return nil, 400, err
	}

	id, err := h.CampaingService.Create(request)

	return map[string]string{"id": id}, 201, err
}

// fiz pra deixar aqui como é o endpoint sem o wrapper
// func (h *Handler) CreateNewCamp(w http.ResponseWriter, r *http.Request) {
// 	var request contract.NewCampaign

// 	err := render.DecodeJSON(r.Body, &request)

// 	if err != nil {
// 		render.Status(r, 400)
// 		render.JSON(w, r, map[string]string{"Erro": err.Error()})
// 		return
// 	}

// 	id, err := h.CampaingService.Create(request)

// 	if err != nil {
// 		if errors.Is(err, internaerrors.ErrInternal) {
// 			render.Status(r, 500)
// 		} else {
// 			render.Status(r, 400)
// 			render.JSON(w, r, map[string]string{"Erro": err.Error()})
// 			return
// 		}
// 	}
// 	render.Status(r, 201)
// 	render.JSON(w, r, map[string]string{"id": id})
// }
