package entities

type Credentials struct {
	Id           string `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	PasswordHash []byte
}
