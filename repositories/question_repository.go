package repositories

import "github.com/MikolajRatajczyk/Langmal-Server/entities"

type QuestionRepository interface {
	FindAll() []entities.Question
}

func NewQuestionRepository() QuestionRepository {
	//	TODO: open DB session and pass it to questionRepositoryImpl
	return &questionRepositoryImpl{}
}

type questionRepositoryImpl struct {
	//	TODO: add `connection *gorm.DB` or similar and use it
}

func (qr *questionRepositoryImpl) FindAll() []entities.Question {
	questions := createQuestions()
	return questions
}

//	TODO: remove and use DB instead
func createQuestions() []entities.Question {
	question1 := entities.Question{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer A",
	}
	question2 := entities.Question{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}
	question3 := entities.Question{
		Title:   "Third question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}
	return []entities.Question{question1, question2, question3}
}
