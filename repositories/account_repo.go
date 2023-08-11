package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"gorm.io/gorm"
)

type AccountRepoInterface interface {
	Create(account models.AccountEntity) bool
	Find(email string) (models.AccountEntity, bool)
	CloseDB()
}

func NewAccountRepo(dbName string) AccountRepoInterface {
	return &accountRepo{
		db: getDb(dbName, models.AccountEntity{}),
	}
}

type accountRepo struct {
	db *gorm.DB
}

func (ar *accountRepo) Create(account models.AccountEntity) bool {
	if err := ar.db.Create(account).Error; err != nil {
		log.Println("Failed to create a new account in DB!")
		return false
	} else {
		return true
	}
}

func (ar *accountRepo) Find(email string) (models.AccountEntity, bool) {
	var account models.AccountEntity
	err := ar.db.First(&account, "email = ?", email).Error
	success := err == nil
	return account, success
}

func (ar *accountRepo) CloseDB() {
	sqlDB, err := ar.db.DB()
	if err != nil {
		panic("Failed to get SQL DB!")
	}
	sqlDB.Close()
}
