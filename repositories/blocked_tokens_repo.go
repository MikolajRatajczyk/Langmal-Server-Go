package repositories

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"gorm.io/gorm"
)

type BlockedTokenRepoInterface interface {
	Add(id string) bool
	IsBlocked(id string) bool
}

func NewBlockedTokenRepo(dbName string) BlockedTokenRepoInterface {
	return &blockedTokenRepo{
		db: getDb(dbName, &models.BlockedTokenEntity{}),
	}
}

type blockedTokenRepo struct {
	db *gorm.DB
}

func (btr *blockedTokenRepo) Add(id string) bool {
	blockedToken := models.BlockedTokenEntity{Id: id}
	err := btr.db.Create(&blockedToken).Error
	return err == nil
}

func (btr *blockedTokenRepo) IsBlocked(id string) bool {
	var blockedToken models.BlockedTokenEntity
	err := btr.db.First(&blockedToken, "id = ?", id).Error
	return err == nil
}
