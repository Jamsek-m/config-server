package services

import (
	"../db"
	"../errors"
	"../models"
	"fmt"
)

type RoleService struct{}

func (r RoleService) GetRoles() ([]models.Role, int, error) {
	roles := make([]models.Role, 0)
	err := db.GetConnection().Table(models.ROLE_TABLE_NAME).Find(&roles).Error
	if err != nil {
		fmt.Println(err)
		return nil, -1, errors.InternalServerError
	}
	return roles, len(roles), nil
}
