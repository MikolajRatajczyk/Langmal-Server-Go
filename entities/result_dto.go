package entities

type ResultDto struct {
	Correct   int    `json:"correct" binding:"required"`
	Total     int    `json:"total" binding:"required"`
	TestId    string `json:"test_id" binding:"required"`
	CreatedAt int64  `json:"created_at"`
}
