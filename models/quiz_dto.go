package models

type QuizDto struct {
	Title     string        `json:"title"`
	Id        string        `json:"id"`
	Questions []QuestionDto `json:"questions"`
}
