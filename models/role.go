package models

const ROLE_TABLE_NAME = "roles"

const ROLE_READ = "CONF_READ"
const ROLE_WRITE = "CONF_WRITE"
const ROLE_ADMIN = "ADMIN"

type Role struct {
	ID    uint   `json:"id,omitempty"`
	Code  string `json:"code,omitempty"`
	Label string `json:"label,omitempty"`
}

func (r Role) TableName() string {
	return ROLE_TABLE_NAME
}
