package repository

import (
	"github.com/REIS0/scuz/database"
	"github.com/REIS0/scuz/models"
)

func GetUser(id string) models.User {
	var user models.User
	database.DBConn.Find(&user, "Username = ? OR Email = ?", id, id)
	return user
}

func CreateUser(user models.User) {
	database.DBConn.Create(&user)
}
