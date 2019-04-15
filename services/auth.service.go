package services

import (
	"fmt"
	"github.com/Jamsek-m/config-server/config"
	"github.com/Jamsek-m/config-server/db"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/random"
	"github.com/Jamsek-m/config-server/response"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type AuthService struct{}

func (a AuthService) Login(req *http.Request, dto models.LoginRequest, res http.ResponseWriter) error {
	userService := UserService{}
	validLogin, user, err := userService.ValidateLogin(dto)
	if err != nil && err == errors.NotFoundError {
		return errors.UnAuthorizedError
	} else if err != nil {
		return err
	} else if !validLogin {
		return errors.UnAuthorizedError
	}

	session, sessionErr := a.startLoginSession(req, user)
	if sessionErr != nil {
		return sessionErr
	}
	isProduction := config.GetConfiguration().Service.Env == "prod"
	cookie := http.Cookie{
		Name:     response.COOKIE_NAME,
		HttpOnly: true,
		Secure:   isProduction,
		MaxAge:   config.GetConfiguration().Server.SessionDuration,
		Path:     "/",
		Value:    session.SessionID,
	}
	http.SetCookie(res, &cookie)
	return nil
}

func (a AuthService) AuthorizeAccess(req *http.Request) (*models.User, error) {
	cookie := a.getSessionCookie(req)
	if cookie == nil {
		return nil, errors.UnAuthorizedError
	}
	session, sessErr := a.getLoginSession(cookie.Value)
	if sessErr != nil {
		return nil, sessErr
	}
	ip := req.RemoteAddr

	if session.Ip != ip {
		return nil, errors.UnAuthorizedError
	}
	if time.Now().After(session.ExpirationDate) {
		return nil, errors.UnAuthorizedError
	}
	userService := UserService{}
	user, _ := userService.GetUserById(session.UserID)
	return user, nil
}

func (a AuthService) getSessionCookie(req *http.Request) *http.Cookie {
	for _, cookie := range req.Cookies() {
		if cookie.Name == response.COOKIE_NAME {
			return cookie
		}
	}
	return nil
}

func (a AuthService) LogoutUser(res http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     response.COOKIE_NAME,
		HttpOnly: true,
		MaxAge:   0,
		Path:     "/",
		Value:    "LOGOUT",
	}
	http.SetCookie(res, &cookie)
}

func (a AuthService) deleteAllOtherSessions(ip string) error {
	err := db.GetConnection().Table(models.SESSION_TABLE_NAME).Delete(&models.Session{}).Where("ip = ?", ip).Error
	if err != nil {
		fmt.Println(err)
		return errors.InternalServerError
	}
	return nil
}

func (a AuthService) startLoginSession(req *http.Request, user *models.User) (*models.Session, error) {
	ip := req.RemoteAddr
	deleteErr := a.deleteAllOtherSessions(ip)
	if deleteErr != nil {
		return nil, deleteErr
	}
	sessionId := random.GenerateRandomString(10)
	now := time.Now()
	expirationDate := now.Add(time.Duration(config.GetConfiguration().Server.SessionDuration) * time.Second)

	session := &models.Session{
		Ip:             ip,
		UserID:         user.ID,
		SessionID:      sessionId,
		ExpirationDate: expirationDate,
	}

	err := db.GetConnection().Create(session).Error

	if err != nil {
		fmt.Println(err)
		return nil, errors.InternalServerError
	} else if session.ID <= 0 {
		return nil, errors.InternalServerError
	}
	return session, nil
}

func (a AuthService) getLoginSession(sessionId string) (*models.Session, error) {
	session := &models.Session{}
	err := db.GetConnection().Table(models.SESSION_TABLE_NAME).Where("session_id = ?", sessionId).First(session).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.NotFoundError
	} else if err != nil {
		fmt.Println(err)
		return nil, errors.InternalServerError
	}
	return session, nil
}
