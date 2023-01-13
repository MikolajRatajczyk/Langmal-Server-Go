package models

type TestDto struct {
	Name      string        `json:"name"`
	Id        string        `json:"id"`
	Questions []QuestionDto `json:"questions"`
}
