package main

import (
	"emailNil/internal/contract"
	"emailNil/internal/domain/campaign"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Route("/campaign", func(r chi.Router) {
		r.Post("/newCampaign", createCampaing)
	})

	http.ListenAndServe(":3000", r)
}

func createCampaing(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign

	if err := render.DecodeJSON(r.Body, &request); err != nil {
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": "JSON inválido"})
		return
	}

	service := campaign.Service{}
	id, err := service.Create(request)

	if err != nil {
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})

}
