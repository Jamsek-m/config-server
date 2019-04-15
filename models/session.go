package models

import (
	"time"
)

const SESSION_TABLE_NAME = "sessions"

type Session struct {
	ID             uint      `json:"-"`
	Ip             string    `json:"-"`
	ExpirationDate time.Time `json:"expirationDate,omitempty" gorm:"column:expiration_date"`
	SessionID      string    `json:"-" gorm:"column:session_id"`
	UserID         uint      `json:"userId,omitempty"`
}

func (s Session) TableName() string {
	return SESSION_TABLE_NAME
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
