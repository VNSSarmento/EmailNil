package campaign

import (
	"emailNil/internal/contract"
	internaerrors "emailNil/internal/internaErrors"
)

//struct service recebe o reporsitorio
//service é como o controller, ele vai ficar em contato com o Repositorio e o Handler

type ServiceImp struct {
	Repository Repository
}

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() (interface{}, error)
	GetId(id string) (*contract.NewCampaignResponse, error)
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internaerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) Get() (interface{}, error) {

	return s.Repository.Get()
}

func (s *ServiceImp) GetId(id string) (*contract.NewCampaignResponse, error) {
	campaign, err := s.Repository.GetId(id)

	if err != nil {
		return nil, internaerrors.ErrInternal
	}

	return &contract.NewCampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
