package repositories

import "github.com/MikolajRatajczyk/Langmal-Server/entities"

//	TODO: It should use Question (not DTO)
type QuestionRepository interface {
	FindAll() []entities.QuestionDto
}

func NewQuestionRepository() QuestionRepository {
	//	TODO: open DB session and pass it to questionRepositoryImpl
	return &questionRepositoryImpl{}
}

type questionRepositoryImpl struct {
	//	TODO: add `connection *gorm.DB` or similar and use it
}

func (qr *questionRepositoryImpl) FindAll() []entities.QuestionDto {
	questions := createQuestions()
	return questions
}

//	TODO: remove and use DB instead
func createQuestions() []entities.QuestionDto {
	question1 := entities.QuestionDto{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer A",
	}
	question2 := entities.QuestionDto{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}
	question3 := entities.QuestionDto{
		Title:   "Third question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}
	return []entities.QuestionDto{question1, question2, question3}
}
