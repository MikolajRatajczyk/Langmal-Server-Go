package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"gorm.io/gorm"
)

type CredentialsRepositoryInterface interface {
	Create(credentials entities.Credentials) bool
	Find(email string) entities.Credentials
	CloseDB()
}

func NewCredentialsRepository() CredentialsRepositoryInterface {
	return &credentialsRepository{
		db: getDb("credentials", entities.Credentials{}),
	}
}

type credentialsRepository struct {
	db *gorm.DB
}

func (cr *credentialsRepository) Create(credentials entities.Credentials) bool {
	if err := cr.db.Create(credentials).Error; err != nil {
		log.Println("Failed to create a new user in DB!")
		return false
	} else {
		return true
	}
}

func (cr *credentialsRepository) Find(email string) entities.Credentials {
	var credentials entities.Credentials
	cr.db.Where("email = ?", email).First(&credentials)
	return credentials
}

func (cr *credentialsRepository) CloseDB() {
	sqlDB, err := cr.db.DB()
	if err != nil {
		panic("Failed to get SQL DB!")
	}
	sqlDB.Close()
}
