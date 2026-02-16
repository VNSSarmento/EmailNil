package main

import (
	"emailNil/internal/domain/campaign"
	"emailNil/internal/endpoints"
	"emailNil/internal/infra/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	CampaignRepository := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaingService: &CampaignRepository,
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/", endpoints.HandlerError(handler.CampaignGet))

	r.Route("/campaign", func(r chi.Router) {
		r.Post("/newCampaign", endpoints.HandlerError(handler.CreateCampaing))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetId))
	})

	http.ListenAndServe(":3000", r)
}
