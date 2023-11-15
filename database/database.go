package database

import (
	"errors"
	"fmt"

	"github.com/REIS0/scuz/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	path string
	conn *gorm.DB
}

func (db *Database) InitDatabase(path string) (*Database, error) {
	var err error
	db.conn, err = gorm.Open(sqlite.Open(db.path), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to start database")
	}
	fmt.Println("Connection to database started")
	db.conn.AutoMigrate(&models.User{})
	fmt.Println("Migrations finished")
	return db, nil
}

func (db *Database) GetConnection() *gorm.DB {
	return db.conn
}

func (db *Database) GetPath() string {
	return db.path
}
