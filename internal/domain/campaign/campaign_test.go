package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campanha"
	content  = "Body"
	contacts = []string{"emailteste@gmail.com", "emailteste2@gmail.com"}
)

func Test_NewCampaign_CreatCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	//assert.Equal(campaign.ID, "d5v2f2s81p6rj2rafsu0")
	assert.Equal(campaign.Name, "Capanha X")

}

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

	_, err := NewCampaign(name, content, contacts)

	assert.Equal("name is required", err.Error())

}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())

}
