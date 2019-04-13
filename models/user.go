package models

const USER_TABLE_NAME = "users"

type User struct {
	ID          uint   `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	AccessToken string `json:"-" gorm:"column:access_token"`
	Password    string `json:"-"`
	Roles       []Role `json:"roles,omitempty" gorm:"many2many:user_roles;"`
}

func (u User) TableName() string {
	return USER_TABLE_NAME
}

type UserRequest struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	PasswordRepeat string `json:"passwordRepeat,omitempty"`
}

type UserResponse struct {
	Username string `json:"username,omitempty"`
}

func NewUserResponse(username string) *UserResponse {
	resp := &UserResponse{}
	resp.Username = username
	return resp
}

type TokenResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

func NewTokenResponse(token string) *TokenResponse {
	resp := &TokenResponse{}
	resp.AccessToken = token
	return resp
}
