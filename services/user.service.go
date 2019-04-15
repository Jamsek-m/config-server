package services

import (
	"fmt"
	"github.com/Jamsek-m/config-server/db"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (u UserService) GetAllUsers() ([]models.User, int, error) {
	users := make([]models.User, 0)
	err := db.GetConnection().Table(models.USER_TABLE_NAME).Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil, -1, errors.InternalServerError
	}
	return users, len(users), nil
}

func (u UserService) GetUserById(id uint) (*models.User, error) {
	user := &models.User{}
	err := db.GetConnection().Table(models.USER_TABLE_NAME).Where("id = ?", id).First(user).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.NotFoundError
	} else if err != nil {
		fmt.Println(err)
		return nil, errors.InternalServerError
	}
	return user, nil
}

func (u UserService) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := db.GetConnection().Table(models.USER_TABLE_NAME).Where("username = ?", username).First(user).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.NotFoundError
	} else if err != nil {
		return nil, errors.InternalServerError
	}
	return user, nil
}

func (u UserService) ValidateLogin(req models.LoginRequest) (bool, *models.User, error) {
	user, err := u.GetUserByUsername(req.Username)
	if err != nil && err == errors.NotFoundError {
		fmt.Println(err)
		return false, nil, errors.UnAuthorizedError
	} else if err != nil {
		fmt.Println(err)
		return false, nil, errors.InternalServerError
	}
	compareError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if compareError != nil {
		return false, nil, errors.UnAuthorizedError
	}
	return true, user, nil
}

func (u UserService) CreateUser(req *models.UserRequest) (*models.UserResponse, error) {

	if req.Username == "" {
		return nil, errors.NilValidationFailedError
	}
	if req.Password != req.PasswordRepeat {
		return nil, errors.SemanticValidationFailedError
	}
	if req.Password == "" {
		return nil, errors.NilValidationFailedError
	}

	user := &models.User{}
	user.Username = req.Username
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		fmt.Println(hashErr)
		return nil, errors.InternalServerError
	}
	user.Password = string(hashedPassword)
	db.GetConnection().Create(user)

	if user.ID <= 0 {
		return nil, errors.InternalServerError
	}

	return models.NewUserResponse(user.Username), nil
}
