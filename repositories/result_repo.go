package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"gorm.io/gorm"
)

type ResultRepoInterface interface {
	Create(result models.Result) bool
	Find(accountId string) ([]models.Result, bool)
}

func NewResultRepo(dbName string) ResultRepoInterface {
	return &resultRepo{
		db: getDb(dbName, models.Result{}),
	}
}

type resultRepo struct {
	db *gorm.DB
}

func (rr *resultRepo) Find(accountId string) ([]models.Result, bool) {
	var results []models.Result
	rr.db.Where("account_id = ?", accountId).Find(&results)
	success := len(results) > 0
	return results, success
}

func (rr *resultRepo) Create(result models.Result) bool {
	err := rr.db.Create(&result).Error
	if err != nil {
		log.Println("Failed to create a new result in DB!")
		return false
	} else {
		return true
	}
}
