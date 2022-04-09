package entity

type HashedCredentials struct {
	Username     string `gorm:"primaryKey"`
	PasswordHash []byte
}
