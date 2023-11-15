package repository

import (
	"github.com/REIS0/scuz/database"
	"github.com/REIS0/scuz/models"
)

type UserRepository struct {
	DB database.Database
}

func (ur *UserRepository) GetUser(id string) models.User {
	var user models.User
	conn := ur.DB.GetConnection()

	conn.Find(&user, "Username = ? OR Email = ?", id, id)
	return user
}

func (ur *UserRepository) CreateUser(user models.User) {
	conn := ur.DB.GetConnection()
	conn.Create(&user)
}
