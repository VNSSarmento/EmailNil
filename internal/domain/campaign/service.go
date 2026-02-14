package campaign

import (
	"emailNil/internal/contract"
	internaerrors "emailNil/internal/internaErrors"
)

//struct service recebe o reporsitorio
//service é como o controller, ele vai ficar em contato com o Repositorio e o Handler

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

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
