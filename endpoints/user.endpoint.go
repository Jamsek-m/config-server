package endpoints

import (
	"encoding/json"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserEndpoint struct{}

func (u UserEndpoint) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	users, usersCount, err := userService.GetAllUsers()
	if err != nil {
		response.HandleError(res, err)
	} else {
		response.Json(res, response.JsonArgs{
			Data:    users,
			Headers: map[string]string{response.X_TOTAL_COUNT: strconv.Itoa(usersCount)},
		})
	}
}

func (u UserEndpoint) GetUserById(res http.ResponseWriter, req *http.Request) {
	userId, paramErr := strconv.Atoi(mux.Vars(req)["id"])
	if paramErr != nil {
		response.HandleError(res, errors.BadRequestError)
		return
	}
	user, err := userService.GetUserById(uint(userId))
	if err != nil {
		response.HandleError(res, err)
	} else {
		response.Json(res, response.JsonArgs{Data: user})
	}
}

func (u UserEndpoint) CreateUser(res http.ResponseWriter, req *http.Request) {
	userRequest := &models.UserRequest{}
	decodeErr := json.NewDecoder(req.Body).Decode(userRequest)
	if decodeErr != nil {
		response.HandleError(res, errors.BadRequestError)
	} else {
		resp, err := userService.CreateUser(userRequest)
		if err != nil {
			response.HandleError(res, err)
		} else {
			response.Json(res, response.JsonArgs{Data: resp, Status: http.StatusCreated})
		}
	}
}
