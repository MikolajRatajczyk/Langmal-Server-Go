package repository

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entity"
)

type QuestionRepository interface {
	FindAll() []entity.Question
}

func NewQuestionRepository() QuestionRepository {
	//	TODO: open DB session and pass it to questionRepositoryImpl
	return &questionRepositoryImpl{}
}

type questionRepositoryImpl struct {
	//	TODO: add `connection *gorm.DB` or similar and use it
}

func (qr *questionRepositoryImpl) FindAll() []entity.Question {
	questions := createQuestions()
	return questions
}

//	TODO: remove and use DB instead
func createQuestions() []entity.Question {
	question1 := entity.Question{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer A",
	}
	question2 := entity.Question{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}
	question3 := entity.Question{
		Title:   "Third question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}
	return []entity.Question{question1, question2, question3}
}
