package models

import "gorm.io/gorm"

type User struct {
	Id    int    `gorm:"primary_key" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null;unique_index" json:"email"`
	gorm.Model
}
