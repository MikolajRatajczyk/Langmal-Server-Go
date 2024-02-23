package models

type ResultReadDto struct {
	ResultWriteDto
	QuizTitle string `json:"quiz_title" binding:"required"`
}
