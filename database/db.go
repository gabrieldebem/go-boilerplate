package database

import (
	"github.com/boilerplate/packages/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var lock = &sync.Mutex{}
var db *gorm.DB

type Database struct{}

// StartDb Start the database connection.
func (d Database) StartDb() {
	dns := d.getDns()

	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	db = database

	if err != nil {
		panic(err)
	}

	d.runMigrations()
}

// getDns Get the database connection string.
func (d Database) getDns() string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	host := "host=" + viper.Get("DB_HOST").(string)
	user := " user=" + viper.Get("DB_USER").(string)
	dbname := " dbname=" + viper.Get("DB_NAME").(string)
	port := " port=" + viper.Get("DB_PORT").(string)
	sslmode := " sslmode=" + viper.Get("DB_SSLMODE").(string)
	password := " password=" + viper.Get("DB_PASSWORD").(string)

	return host + user + password + dbname + port + sslmode
}

// runMigrations Run the database migrations.
func (d Database) runMigrations() {
	err := db.AutoMigrate(
		&models.User{},
		&models.Log{},
	)

	if err != nil {
		panic(err)
	}
}

// GetDb Get the database connection.
func (d Database) GetDb() *gorm.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		if db == nil {
			d.StartDb()
		}
	}

	return db
}
