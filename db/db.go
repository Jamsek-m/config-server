package db

import (
	"fmt"
	"github.com/Jamsek-m/config-server/config"
	"github.com/Jamsek-m/config-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB

func ConnectToDatabase() {
	connection, connectionErr := gorm.Open(
		config.GetConfiguration().Datasource.Type,
		config.GetConfiguration().Datasource.Location)

	if connectionErr != nil {
		fmt.Printf("DB connection error! " + connectionErr.Error())
	}
	fmt.Printf("Connected to database %s!\n", config.GetConfiguration().Datasource.Location)
	conn = connection
	conn.Debug().AutoMigrate(&models.User{}, &models.Role{}, &models.Config{}, &models.Session{})
}

func GetConnection() *gorm.DB {
	return conn
}
