package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CredentialsRepositoryInterface interface {
	Create(credentials entities.HashedCredentials) bool
	Find(username string) entities.HashedCredentials
	CloseDB()
}

func NewCredentialsRepository() CredentialsRepositoryInterface {
	return &credentialsRepository{
		db: getDb(),
	}
}

type credentialsRepository struct {
	db *gorm.DB
}

func (cr *credentialsRepository) Create(credentials entities.HashedCredentials) bool {
	if err := cr.db.Create(credentials).Error; err != nil {
		log.Println("Failed to create a new user in DB!")
		return false
	} else {
		return true
	}
}

func (cr *credentialsRepository) Find(username string) entities.HashedCredentials {
	var hashedCredentials entities.HashedCredentials
	cr.db.Where("username = ?", username).First(&hashedCredentials)
	return hashedCredentials
}

func (cr *credentialsRepository) CloseDB() {
	sqlDB, err := cr.db.DB()
	if err != nil {
		panic("Failed to get SQL DB!")
	}
	sqlDB.Close()
}

func getDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("credentials.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(entities.HashedCredentials{})
	return db
}