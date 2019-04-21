package endpoints

import (
	"encoding/json"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/response"
	"net/http"
	"strings"
)

type ConfigEndpoint struct{}

func (c ConfigEndpoint) GetConfigByKey(res http.ResponseWriter, req *http.Request) {
	_, validToken, tokenErr := tokenService.ValidateToken(req)
	if !validToken {
		response.HandleError(res, tokenErr)
		return
	}

	configKey := strings.Replace(req.URL.Path, "/v1/keys", "", 1)

	configEntry, err := configService.GetConfigByKey(configKey)
	if err != nil {
		response.HandleError(res, err)
	} else {
		response.Json(res, response.JsonArgs{Data: configEntry})
	}
}

func (c ConfigEndpoint) CreateConfigEntry(res http.ResponseWriter, req *http.Request) {
	_, validToken, tokenErr := tokenService.ValidateToken(req)
	if !validToken {
		response.HandleError(res, tokenErr)
		return
	}

	configKey := strings.Replace(req.URL.Path, "/v1/keys", "", 1)
	requestBody, decodeErr := decodeConfigRequest(req)
	if decodeErr != nil {
		response.HandleError(res, errors.BadRequestError)
	} else {
		requestBody.Key = configKey
		responseBody, err := configService.CreateOrUpdateConfigEntry(requestBody)
		if err != nil {
			response.HandleError(res, err)
		} else {
			response.Json(res, response.JsonArgs{Data: responseBody, Status: http.StatusCreated})
		}
	}
}

func decodeConfigRequest(req *http.Request) (*models.ConfigRequest, error) {
	requestBody := &models.ConfigRequest{}
	decodeErr := json.NewDecoder(req.Body).Decode(requestBody)

	if decodeErr != nil {
		return nil, decodeErr
	}
	return requestBody, nil
}
