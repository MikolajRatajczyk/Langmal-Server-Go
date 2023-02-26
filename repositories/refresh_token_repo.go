package repositories

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepoInterface interface {
	Create(tokenHash []byte, accountId string, deviceId string) bool
	Find(accountId string) ([]models.AssociatedToken, bool)
}

func NewRefreshTokenRepo(dbName string) RefreshTokenRepoInterface {
	return &refreshTokenRepo{
		db: getDb(dbName, models.AssociatedToken{}),
	}
}

type refreshTokenRepo struct {
	db *gorm.DB
}

func (rr *refreshTokenRepo) Create(tokenHash []byte, accountId string, deviceId string) bool {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return false
	}

	associatedToken := models.AssociatedToken{
		Uuid:      uuid.String(),
		AccountId: accountId,
		DeviceId:  deviceId,
		TokenHash: tokenHash,
	}

	err = rr.db.Create(associatedToken).Error
	if err != nil {
		log.Println("Failed to create a new associated refresh token in DB!")
		return false
	}

	return true
}

func (rr *refreshTokenRepo) Find(accountId string) ([]models.AssociatedToken, bool) {
	var associatedTokens []models.AssociatedToken
	result := rr.db.Where("account_id = ?", accountId).Find(&associatedTokens)

	if result.Error != nil {
		return []models.AssociatedToken{}, false
	}
	return associatedTokens, true
}
