package endpoints

import "emailNil/internal/domain/campaign"

type Handler struct {
	CampaingService campaign.Service
}
