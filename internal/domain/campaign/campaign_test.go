package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreatCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Capanha X"
	content := "Body"
	contacts := []string{"emailteste@gmail.com", "emailteste2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, 10)
	assert.Equal(campaign.Name, "Capanha X")

}

func Test_NewCampaign_IDisNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Capanha X"
	content := "Body"
	contacts := []string{"emailteste@gmail.com", "emailteste2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

func Test_NewCampaign_CreatAtNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Capanha X"
	content := "Body"
	contacts := []string{"emailteste@gmail.com", "emailteste2@gmail.com"}
	now := time.Now().Add(-time.Hour)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.CreatAt)
	//Maior que a variavel Now
	assert.Greater(campaign.CreatAt, now)

}
