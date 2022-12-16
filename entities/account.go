package entities

type Account struct {
	Id           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash []byte
}
