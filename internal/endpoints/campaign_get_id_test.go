package endpoints

import (
	"emailNil/internal/contract"
	serviceMock "emailNil/internal/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Campaing_GetId(t *testing.T) {
	assert := assert.New(t)

	NewCampaignResponse := contract.NewCampaignResponse{
		ID:      xid.New().String(),
		Name:    "Test tttt",
		Content: "Bodsssy",
		Status:  "Pending",
	}
	service := new(serviceMock.CampaignServiceMock)

	handler := Handler{
		CampaingService: service,
	}

	service.On("GetId", mock.Anything).Return(&NewCampaignResponse, nil)

	request, _ := http.NewRequest("GET", "/", nil)

	response := httptest.NewRecorder()

	responseId, status, _ := handler.CampaignGetId(response, request)

	assert.Equal(200, status)
	assert.Equal(NewCampaignResponse.ID, responseId.(*contract.NewCampaignResponse).ID)

}
