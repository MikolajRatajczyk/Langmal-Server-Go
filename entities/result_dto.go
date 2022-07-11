package entities

type ResultDto struct {
	Correct   int    `json:"correct" binding:"required"`
	Wrong     int    `json:"wrong" binding:"required"`
	TestId    string `json:"test_id" binding:"required"`
	CreatedAt int64  `json:"created_at"`
}
