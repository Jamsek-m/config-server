package services

import (
	"github.com/Jamsek-m/config-server/config"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/response"
	"net/http"
)

type AuthService struct{}

func (a AuthService) Login(req models.LoginRequest, res http.ResponseWriter) {

}

func (a AuthService) Logout(res http.ResponseWriter) {

}

func (a AuthService) GetCurrentUser() *models.User {
	return nil
}

func (a AuthService) LogoutUser(res http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     response.COOKIE_NAME,
		HttpOnly: true,
		MaxAge:   config.GetConfiguration().Server.SessionDuration,
		Path:     "/",
		Value:    "",
	}
	http.SetCookie(res, &cookie)
}

func (a AuthService) deleteAllOtherSessions(ip string) {

}

func (a AuthService) startLoginSession(user models.User) *models.Session {
	return nil
}

func (a AuthService) getLoginSession(ip string) *models.Session {
	return nil
}
