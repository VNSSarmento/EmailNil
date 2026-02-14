package endpoints

import (
	internaerrors "emailNil/internal/internaErrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_internal_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internaerrors.ErrInternal
	}

	handlerFunc := HandlerError(endpoint)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	handlerFunc.ServeHTTP(response, request)

	assert.Equal(http.StatusInternalServerError, response.Code)
	assert.Contains(response.Body.String(), internaerrors.ErrInternal.Error())
}

func Test_HandlerError_when_domain_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("error domain")
	}

	handlerFunc := HandlerError(endpoint)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	handlerFunc.ServeHTTP(response, request)

	assert.Equal(http.StatusBadRequest, response.Code)
	assert.Contains(response.Body.String(), "error domain")
}

func Test_HandlerError_when_obj_status(t *testing.T) {
	assert := assert.New(t)

	type obj struct {
		Id string
	}

	objExpected := obj{
		Id: "sklldl24sfUa32",
	}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}

	handlerFunc := HandlerError(endpoint)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	handlerFunc.ServeHTTP(response, request)

	assert.Equal(http.StatusCreated, response.Code)
	objInstaced := obj{}
	json.Unmarshal(response.Body.Bytes(), &objInstaced)
	assert.Equal(objExpected, objInstaced)

}
