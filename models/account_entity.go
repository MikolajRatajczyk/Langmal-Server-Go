package models

type AccountEntity struct {
	Id           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash []byte
}
