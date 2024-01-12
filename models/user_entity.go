package models

type UserEntity struct {
	Id           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash []byte
}
