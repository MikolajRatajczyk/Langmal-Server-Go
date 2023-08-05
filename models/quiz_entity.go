package models

type QuizEntity struct {
	Title     string
	Id        string
	Questions []QuestionEntity
}
