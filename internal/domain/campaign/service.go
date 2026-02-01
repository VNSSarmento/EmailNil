package campaign

import "emailNil/internal/contract"

//struct service recebe o reporsitorio

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return nil
}
