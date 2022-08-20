package models

import "gorm.io/gorm"

type Log struct {
	UserAgent string `gorm:"not null" json:"user_agent"`
	gorm.Model
}
