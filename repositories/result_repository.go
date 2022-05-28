package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"gorm.io/gorm"
)

type ResultRepositoryInterface interface {
	Create(result entities.Result) bool
	Find(username string) []entities.Result
}

func NewResultRepository() ResultRepositoryInterface {
	return &resultRepository{
		db: getDb("results", entities.Result{}),
	}
}

type resultRepository struct {
	db *gorm.DB
}

func (rr *resultRepository) Find(username string) []entities.Result {
	var results []entities.Result
	rr.db.Where("username = ?", username).Find(&results)
	return results
}

func (rr *resultRepository) Create(result entities.Result) bool {
	err := rr.db.Create(result).Error
	if err != nil {
		log.Println("Failed to create a new result in DB!")
		return false
	} else {
		return true
	}
}
