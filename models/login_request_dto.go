package models

type LoginRequestDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	DeviceId string `json:"device_id" binding:"required"`
}
