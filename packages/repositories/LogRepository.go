package repositories

import (
	"github.com/boilerplate/packages/models"
)

type LogRepository struct {
	Repository
}

func (l LogRepository) Create(userAgent string) (log models.Log, err error) {
	log = models.Log{
		UserAgent: userAgent,
	}

	err = db.Create(&log).Error

	return
}

func (l LogRepository) List() (logs []models.Log, err error) {
	err = db.Find(&logs).Error

	return
}
