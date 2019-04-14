package endpoints

import (
	"github.com/Jamsek-m/config-server/response"
	"net/http"
	"strconv"
)

type RoleEndpoint struct{}

func (r RoleEndpoint) GetRoles(res http.ResponseWriter, req *http.Request) {
	roles, rolesCount, err := roleService.GetRoles()
	if err != nil {
		response.HandleError(res, err)
	} else {
		response.Json(res, response.JsonArgs{
			Data:    roles,
			Headers: map[string]string{response.X_TOTAL_COUNT: strconv.Itoa(rolesCount)},
		})
	}
}
