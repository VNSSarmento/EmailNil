package database

import (
	"emailNil/internal/domain/campaign"
	internaerrors "emailNil/internal/internaErrors"
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

func (c *CampaignRepository) GetId(id string) (*campaign.Campaign, error) {
	for _, campaign := range c.Campaign {
		if campaign.ID == id {
			return &campaign, nil
		}
	}

	err := internaerrors.ErrInternal

	return nil, err
}
