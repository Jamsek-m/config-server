package models

const ROLE_TABLE_NAME = "roles"

type Role struct {
	ID    uint   `json:"id,omitempty"`
	Code  string `json:"code,omitempty"`
	Label string `json:"label,omitempty"`
}

func (r Role) TableName() string {
	return ROLE_TABLE_NAME
}
