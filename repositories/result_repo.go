package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"gorm.io/gorm"
)

type ResultRepoInterface interface {
	Create(result entities.Result) bool
	Find(accountId string) []entities.Result
}

func NewResultRepo() ResultRepoInterface {
	return &resultRepo{
		db: getDb("results", entities.Result{}),
	}
}

type resultRepo struct {
	db *gorm.DB
}

func (rr *resultRepo) Find(accountId string) []entities.Result {
	var results []entities.Result
	rr.db.Where("account_id = ?", accountId).Find(&results)
	return results
}

func (rr *resultRepo) Create(result entities.Result) bool {
	err := rr.db.Create(&result).Error
	if err != nil {
		log.Println("Failed to create a new result in DB!")
		return false
	} else {
		return true
	}
}
