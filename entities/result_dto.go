package entities

type ResultDto struct {
	Correct   int    `json:"correct"`
	Total     int    `json:"total"`
	TestId    string `json:"test_id"`
	CreatedAt int64  `json:"created_at"`
}
