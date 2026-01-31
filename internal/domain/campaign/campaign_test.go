package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Capanha X"
	content := "Body"
	contacts := []string{"emailteste@gmail.com", "emailteste2@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, 10)
	assert.Equal(campaign.Name, "Capanha X")
}
