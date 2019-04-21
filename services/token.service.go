package services

import (
	"fmt"
	"github.com/Jamsek-m/config-server/db"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/Jamsek-m/config-server/random"
	"github.com/Jamsek-m/config-server/response"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type TokenService struct{}

func (t TokenService) ValidateToken(req *http.Request) (*models.Token, bool, error) {
	rawToken := req.Header.Get(response.X_TOKEN)
	method := req.Method

	var token models.Token
	tokenIndex := strings.Split(rawToken, "-")[0]
	err := db.GetConnection().Table(models.TOKEN_TABLE_NAME).Where("token_index = ?", tokenIndex).First(&token).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, false, errors.UnAuthorizedError
	} else if err != nil {
		fmt.Println(err)
		return nil, false, errors.InternalServerError
	}

	userService := UserService{}
	user, err := userService.GetUserById(token.UserID)
	if err != nil {
		fmt.Println(err)
		return nil, false, errors.InternalServerError
	}

	if method == "" || method == http.MethodGet {
		if !user.HasRole(models.ROLE_READ) {
			return nil, false, errors.ForbiddenError
		}
	} else {
		if !user.HasRole(models.ROLE_WRITE) {
			return nil, false, errors.ForbiddenError
		}
	}

	return &token, true, nil
}

func (t TokenService) getIndexCandidate(index string) bool {
	token := &models.Token{}
	err := db.GetConnection().Table(models.TOKEN_TABLE_NAME).Where("token_index = ?", index).First(token).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return true
	} else if err != nil {
		fmt.Println(err)
		return false
	}
	return false
}

func (t TokenService) generateTokenIndex() string {
	indexCandidate := random.GenerateRandomString(8)
	ok := t.getIndexCandidate(indexCandidate)
	for !ok {
		indexCandidate = random.GenerateRandomString(8)
		ok = t.getIndexCandidate(indexCandidate)
	}
	return indexCandidate
}

func (t TokenService) ExpireToken(tokenId uint) error {
	return db.GetConnection().Where("id = ?", tokenId).Delete(models.Token{}).Error
}

func (t TokenService) GenerateNewToken(req *models.TokenRequest, currentUser *models.User) (*models.TokenResponse, error) {
	token := &models.Token{}
	token.Name = req.Name
	token.UserID = currentUser.ID

	tokenIndex := t.generateTokenIndex()
	token.Index = tokenIndex

	rawToken := token.Index + "-" + random.GenerateRandomString(20)

	hashedToken, hashErr := bcrypt.GenerateFromPassword([]byte(rawToken), bcrypt.DefaultCost)
	if hashErr != nil {
		fmt.Println(hashErr)
		return nil, errors.InternalServerError
	}

	token.Value = string(hashedToken)

	db.GetConnection().Create(token)

	if token.ID <= 0 {
		fmt.Println("Uknown error! Id is 0!")
		return nil, errors.InternalServerError
	}

	resp := models.NewTokenResponse(req.Name, rawToken)
	return &resp, nil
}
