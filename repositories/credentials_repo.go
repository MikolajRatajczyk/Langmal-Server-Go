package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"gorm.io/gorm"
)

type CredentialsRepoInterface interface {
	Create(credentials entities.Credentials) bool
	Find(email string) entities.Credentials
	CloseDB()
}

func NewCredentialsRepo() CredentialsRepoInterface {
	return &credentialsRepo{
		db: getDb("credentials", entities.Credentials{}),
	}
}

type credentialsRepo struct {
	db *gorm.DB
}

func (cr *credentialsRepo) Create(credentials entities.Credentials) bool {
	if err := cr.db.Create(credentials).Error; err != nil {
		log.Println("Failed to create a new user in DB!")
		return false
	} else {
		return true
	}
}

func (cr *credentialsRepo) Find(email string) entities.Credentials {
	var credentials entities.Credentials
	cr.db.Where("email = ?", email).First(&credentials)
	return credentials
}

func (cr *credentialsRepo) CloseDB() {
	sqlDB, err := cr.db.DB()
	if err != nil {
		panic("Failed to get SQL DB!")
	}
	sqlDB.Close()
}
