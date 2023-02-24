package models

type AssociatedToken struct {
	Uuid      string `gorm:"primaryKey"`
	AccountId string
	DeviceId  string
	TokenHash []byte
}
