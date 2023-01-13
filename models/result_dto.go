package models

type ResultDto struct {
	Correct   int    `json:"correct" binding:"number"`
	Wrong     int    `json:"wrong" binding:"number"`
	TestId    string `json:"test_id" binding:"required"`
	CreatedAt int64  `json:"created_at"`
}
