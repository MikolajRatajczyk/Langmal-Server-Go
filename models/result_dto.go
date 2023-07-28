package models

type ResultDtoSave struct {
	Correct   int    `json:"correct" binding:"number"`
	Wrong     int    `json:"wrong" binding:"number"`
	QuizId    string `json:"quiz_id" binding:"required"`
	CreatedAt int64  `json:"created_at" binding:"required"`
}

type ResultDtoRead struct {
	ResultDtoSave
	QuizTitle string `json:"quiz_title" binding:"required"`
}
