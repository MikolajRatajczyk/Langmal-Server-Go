package models

type QuizDto struct {
	Name      string        `json:"name"`
	Id        string        `json:"id"`
	Questions []QuestionDto `json:"questions"`
}
