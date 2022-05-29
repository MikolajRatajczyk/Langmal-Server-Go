package entities

type HashedCredentials struct {
	Id           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash []byte
}
