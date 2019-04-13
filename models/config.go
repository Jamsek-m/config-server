package models

const CONFIG_TABLE_NAME = "configuration"

type Config struct {
	ID    uint   `json:"id,omitempty"`
	Key   string `json:"key,omitempty" gorm:"column:config_key;"`
	Value string `json:"value,omitempty" gorm:"column:config_value;"`
}

func (c Config) TableName() string {
	return CONFIG_TABLE_NAME
}

type ConfigRequest struct {
	Key   string `json:",omitempty"`
	Value string `json:"value,omitempty"`
}

type ConfigResponse struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
