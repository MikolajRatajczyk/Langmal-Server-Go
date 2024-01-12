package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Create(user models.UserEntity) bool
	Find(email string) (models.UserEntity, bool)
}

func NewUserRepo(dbName string) UserRepoInterface {
	return &userRepo{
		db: getDb(dbName, &models.UserEntity{}),
	}
}

type userRepo struct {
	db *gorm.DB
}

func (ar *userRepo) Create(user models.UserEntity) bool {
	if err := ar.db.Create(&user).Error; err != nil {
		log.Println("Failed to create a new user in DB!")
		return false
	} else {
		return true
	}
}

func (ar *userRepo) Find(email string) (models.UserEntity, bool) {
	var user models.UserEntity
	err := ar.db.First(&user, "email = ?", email).Error
	success := err == nil
	return user, success
}
