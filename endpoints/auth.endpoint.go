package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/response"
	"net/http"
)

type AuthEndpoint struct{}

func (a AuthEndpoint) Login(res http.ResponseWriter, req *http.Request) {
	dto := &models.LoginRequest{}
	decodeErr := json.NewDecoder(req.Body).Decode(dto)
	if decodeErr != nil {
		fmt.Println(decodeErr)
		response.HandleError(res, errors.BadRequestError)
	} else {
		authErr := authService.Login(req, *dto, res)
		if authErr != nil {
			response.HandleError(res, authErr)
		} else {
			response.Json(res, response.JsonArgs{Headers: map[string]string{
				response.ACCESS_CONTROL_ALLOW_CREDENTIALS: "true",
			}})
		}
	}
}

func (a AuthEndpoint) Logout(res http.ResponseWriter, req *http.Request) {
	authService.LogoutUser(res)
	response.Json(res, response.JsonArgs{Headers: map[string]string{
		response.ACCESS_CONTROL_ALLOW_CREDENTIALS: "true",
	}})
}

func (a AuthEndpoint) IsAuthorized(res http.ResponseWriter, req *http.Request) {
	_, authErr := authService.AuthorizeAccess(req)
	if authErr != nil {
		response.HandleError(res, errors.UnAuthorizedError)
	} else {
		response.Json(res, response.JsonArgs{})
	}
}

func (a AuthEndpoint) GetCurrentUser(res http.ResponseWriter, req *http.Request) {
	user, authErr := authService.AuthorizeAccess(req)
	if authErr != nil {
		response.HandleError(res, errors.UnAuthorizedError)
	} else {
		response.Json(res, response.JsonArgs{Data: user})
	}
}
