package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"gorm.io/gorm"
)

type ResultRepoInterface interface {
	Create(result models.ResultEntity) bool
	Find(accountId string) []models.ResultEntity
}

func NewResultRepo(dbName string) ResultRepoInterface {
	return &resultRepo{
		db: getDb(dbName, models.ResultEntity{}),
	}
}

type resultRepo struct {
	db *gorm.DB
}

func (rr *resultRepo) Find(accountId string) []models.ResultEntity {
	var results []models.ResultEntity
	rr.db.Find(&results, "account_id = ?", accountId)
	return results
}

func (rr *resultRepo) Create(result models.ResultEntity) bool {
	err := rr.db.Create(&result).Error
	if err != nil {
		log.Println("Failed to create a new result in DB!")
		return false
	} else {
		return true
	}
}
