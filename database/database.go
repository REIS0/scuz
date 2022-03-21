package database

import (
	"errors"
	"fmt"

	"github.com/REIS0/scuz/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase(path string) error {
	var err error
	DBConn, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return errors.New("failed to start database")
	}
	fmt.Println("Connection to database started")
	DBConn.AutoMigrate(&models.User{})
	fmt.Println("Migrations finished")
	return nil
}
