package endpoints

import (
	"encoding/json"
	"github.com/Jamsek-m/config-server/auth"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TokenEndpoint struct{}

func (t TokenEndpoint) CreateToken(res http.ResponseWriter, req *http.Request) {

	currentUser, authenticated, authErr := auth.GetAuthContext(req)
	if !authenticated {
		response.HandleError(res, authErr)
		return
	}

	tokenRequest := &models.TokenRequest{}
	decodeErr := json.NewDecoder(req.Body).Decode(tokenRequest)
	if decodeErr != nil {
		response.HandleError(res, errors.BadRequestError)
	} else {
		resp, err := tokenService.GenerateNewToken(tokenRequest, currentUser)
		if err != nil {
			response.HandleError(res, err)
		} else {
			response.Json(res, response.JsonArgs{Data: resp, Status: http.StatusCreated})
		}
	}
}

func (t TokenEndpoint) ExpireToken(res http.ResponseWriter, req *http.Request) {
	tokenId, paramErr := strconv.Atoi(mux.Vars(req)["id"])
	if paramErr != nil {
		response.HandleError(res, errors.BadRequestError)
		return
	}
	err := tokenService.ExpireToken(uint(tokenId))
	if err != nil {
		response.HandleError(res, err)
	} else {
		response.Json(res, response.JsonArgs{Status: http.StatusNoContent})
	}
}
