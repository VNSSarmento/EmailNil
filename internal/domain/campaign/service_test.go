package campaign

import (
	"emailNil/internal/contract"
	internaerrors "emailNil/internal/internaErrors"
	"errors"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		ID:      xid.New().String(),
		Name:    "Test tttt",
		Content: "Bodsssy",
		Emails:  []string{"al@gmail.com"},
	}

	NewCampaignResponse = contract.NewCampaignResponse{
		ID:      xid.New().String(),
		Name:    "Test tttt",
		Content: "Bodsssy",
		Status:  "Pending",
	}

	fake = faker.New()

	service = ServiceImp{}
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *Campaign) error {
	arrgs := r.Called(campaign)
	return arrgs.Error(0)
}

func (r *RepositoryMock) Get() ([]Campaign, error) {
	//arrgs := r.Called(campaign)
	return nil, nil
}

func (r *RepositoryMock) GetId(id string) (*Campaign, error) {
	arrgs := r.Called(id)
	return arrgs.Get(0).(*Campaign), arrgs.Error(1)
}

func Test_Creat_ValidateDomain(t *testing.T) {

	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internaerrors.ErrInternal, err))

}

func Test_CreatCampaign(t *testing.T) {

	assert := assert.New(t)

	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.NotNil(newCampaign.ID)
	assert.Nil(err)

}

func Test_Creat_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("Erro interno"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internaerrors.ErrInternal, err))
}

func Test_Creat_SaveCampaign(t *testing.T) {

	repositoryMock := new(RepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(Campaign *Campaign) bool {
		if Campaign.Name != newCampaign.Name ||
			Campaign.Content != newCampaign.Content ||
			len(newCampaign.Emails) != len(Campaign.Contacts) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Campaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("asdasdalhlhlhiljhliholholkhiklsd", newCampaign.Content, newCampaign.Emails)

	assert.Equal("Name precisa ter no máximo 24", err.Error())
}

func Test_Campaign_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := &RepositoryMock{}

	service := ServiceImp{
		Repository: repositoryMock,
	}

	repositoryMock.
		On("Save", mock.Anything).
		Return(errors.New("Erro Interno"))

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internaerrors.ErrInternal))
}

func Test_NewCampaignContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(newCampaign.Name, fake.Lorem().Text(1040), newCampaign.Emails)

	assert.Equal("Content precisa ter no máximo 1024", err.Error())

}

func Test_NewCampaignEmailInvalid(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(newCampaign.Name, newCampaign.Content, []string{"Emai-invalid"})

	assert.Equal("Email precisa ser um email válido", err.Error())

}

func Test_Campaign_GetID_ReturnCampaig(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign("Valid Name", "Valid Content", []string{"test@email.com"})

	repositoryMock := &RepositoryMock{}

	service := ServiceImp{
		Repository: repositoryMock,
	}

	repositoryMock.
		On("GetId", mock.Anything).
		Return(campaign, nil)

	campaignReturn, _ := service.GetId(campaign.ID)

	assert.Equal(campaign.ID, campaignReturn.ID)
}
