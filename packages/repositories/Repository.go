package repositories

import (
	"github.com/boilerplate/database"
	"gorm.io/gorm"
)

type Repository struct{}

var db *gorm.DB = database.Database{}.GetDb()
