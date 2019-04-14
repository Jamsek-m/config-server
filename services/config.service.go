package services

import (
	"fmt"
	"github.com/Jamsek-m/config-server/db"
	"github.com/Jamsek-m/config-server/errors"
	"github.com/Jamsek-m/config-server/models"
	"github.com/jinzhu/gorm"
)

type ConfigService struct{}

// Returns configuration for given key and service
func (c ConfigService) GetConfigByKey(key string) (*models.Config, error) {
	configEntry := &models.Config{}
	err := db.
		GetConnection().
		Table(models.CONFIG_TABLE_NAME).
		Where("config_key = ?", key).
		First(configEntry).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		fmt.Println("Error " + err.Error())
		return nil, errors.NotFoundError
	} else if err != nil {
		fmt.Println("Error " + err.Error())
		return nil, errors.InternalServerError
	}
	return configEntry, nil
}

func (c ConfigService) CreateOrUpdateConfigEntry(req *models.ConfigRequest) (*models.ConfigResponse, error) {

	configEntry, entryErr := c.GetConfigByKey(req.Key)
	if entryErr != nil && entryErr != errors.NotFoundError {
		return nil, errors.InternalServerError
	} else if entryErr != nil && entryErr == errors.NotFoundError {
		config := &models.Config{}
		config.Key = req.Key
		config.Value = req.Value
		db.GetConnection().Create(config)

		if config.ID <= 0 {
			return nil, errors.InternalServerError
		}
		resp := models.ConfigResponse{}
		resp.Value = config.Value
		resp.Key = config.Key
		return &resp, nil
	} else {
		return c.updateConfigEntry(req, configEntry)
	}
}

func (c ConfigService) updateConfigEntry(req *models.ConfigRequest, entity *models.Config) (*models.ConfigResponse, error) {
	entity.Value = req.Value
	err := db.GetConnection().Save(&entity).Error
	if err != nil {
		fmt.Println(err)
		return nil, errors.InternalServerError
	}
	resp := models.ConfigResponse{}
	resp.Value = entity.Value
	resp.Key = entity.Key
	return &resp, nil
}
