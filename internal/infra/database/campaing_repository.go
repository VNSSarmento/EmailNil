package database

import (
	"emailNil/internal/domain/campaign"
)

type CampaignRepository struct {
	Campaign []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.Campaign = append(c.Campaign, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.Campaign, nil
}
