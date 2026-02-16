package mock

import (
	"emailNil/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	arrgs := r.Called(newCampaign)
	return arrgs.String(0), arrgs.Error(1)
}

func (m *CampaignServiceMock) Get() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func (m *CampaignServiceMock) GetId(id string) (*contract.NewCampaignResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*contract.NewCampaignResponse), args.Error(1)
}
