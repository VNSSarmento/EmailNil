package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// esse é teste de unidades
var (
	name     = "Campanha"
	content  = "Bodaasdy"
	contacts = []string{"emailteste@gmail.com", "emailteste2@gmail.com"}
)

func Test_NewCampaign_IDisNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

func Test_NewCampaign_CreatAtNotNil(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Hour)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreatAt)
	//Maior que a variavel Now
	assert.Greater(campaign.CreatAt, now)

}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("Name precisa ter no mínimo 5", err.Error())

}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("Content precisa ter no mínimo 5", err.Error())
}

func Test_NewCampaign_MustValidateStatus(t *testing.T) {
	assert := assert.New(t)

	Campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal("Pending", Campaign.Status)
}
