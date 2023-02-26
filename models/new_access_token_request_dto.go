package models

type NewAccessTokenRequestDto struct {
	RefreshJwt string `json:"refresh_jwt" binding:"required"`
}
