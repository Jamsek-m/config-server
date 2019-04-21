package models

const TOKEN_TABLE_NAME = "tokens"

type Token struct {
	ID     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty" gorm:"column:name;"`
	Index  string `json:"-" gorm:"column:token_index;unique_index"`
	Value  string `json:"-" gorm:"column:token_value;"`
	UserID uint   `json:""`
}

func (t Token) TableName() string {
	return TOKEN_TABLE_NAME
}

type TokenRequest struct {
	Name string `json:"name,omitempty"`
}

type TokenResponse struct {
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}

func NewTokenResponse(name string, rawToken string) TokenResponse {
	token := TokenResponse{}
	token.Name = name
	token.Token = rawToken
	return token
}
