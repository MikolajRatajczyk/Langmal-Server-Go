package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"gorm.io/gorm"
)

type AccountRepoInterface interface {
	Create(account entities.Account) bool
	Find(email string) entities.Account
	CloseDB()
}

func NewAccountRepo() AccountRepoInterface {
	return &accountRepo{
		db: getDb("accounts", entities.Account{}),
	}
}

type accountRepo struct {
	db *gorm.DB
}

func (ar *accountRepo) Create(account entities.Account) bool {
	if err := ar.db.Create(account).Error; err != nil {
		log.Println("Failed to create a new account in DB!")
		return false
	} else {
		return true
	}
}

func (ar *accountRepo) Find(email string) entities.Account {
	var account entities.Account
	ar.db.Where("email = ?", email).First(&account)
	return account
}

func (ar *accountRepo) CloseDB() {
	sqlDB, err := ar.db.DB()
	if err != nil {
		panic("Failed to get SQL DB!")
	}
	sqlDB.Close()
}
