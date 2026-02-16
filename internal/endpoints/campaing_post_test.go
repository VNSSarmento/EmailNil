package endpoints

import (
	"bytes"
	"emailNil/internal/contract"
	serviceMock "emailNil/internal/test/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaingPost_should_save_new_campaing(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaign{
		Name:    "teste",
		Content: "jis kakop akkf",
		Emails:  []string{"emaiks@gmail.com", "emaiol2@gmail.com"},
	}

	service := new(serviceMock.CampaignServiceMock)

	handler := Handler{
		CampaingService: service,
	}

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name {
			return true
		} else {
			return false
		}
	})).Return("asdasdsad", nil)

	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(body)

	request, _ := http.NewRequest("POST", "/", &buf)

	resonse := httptest.NewRecorder()

	_, status, err := handler.CreateCampaing(resonse, request)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaingPost_should_save_new_campaing_error(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaign{
		Name:    "a",
		Content: "jis kakop akkf",
		Emails:  []string{"emaiks@gmail.com", "emaiol2@gmail.com"},
	}

	service := new(serviceMock.CampaignServiceMock)

	handler := Handler{
		CampaingService: service,
	}

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(body)

	request, _ := http.NewRequest("POST", "/", &buf)

	resonse := httptest.NewRecorder()

	_, _, err := handler.CreateCampaing(resonse, request)

	assert.NotNil(err)

}
