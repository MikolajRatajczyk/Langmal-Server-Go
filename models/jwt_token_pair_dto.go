package models

type JwtTokenPairDto struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}
