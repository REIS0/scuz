package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"primaryKey;foreignKey:UserRefer"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"` //todo: mover para hash depois

}

func (u User) CheckPassword(password string) bool {
	// TODO: converter de hash para comparacao
	return u.Password == password
}
