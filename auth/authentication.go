package auth

import (
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/services"
	"net/http"
)

func GetAuthContext(request *http.Request) (*models.User, bool, error) {
	authService := services.AuthService{}

	user, authErr := authService.AuthorizeAccess(request)
	if authErr != nil {
		return nil, false, errors.UnAuthorizedError
	} else {
		return user, true, nil
	}
}
