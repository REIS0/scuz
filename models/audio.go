package models

import "gorm.io/gorm"

type AudioFile struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserRefer   string
}
