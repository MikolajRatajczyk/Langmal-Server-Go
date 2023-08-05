package repositories

import "github.com/MikolajRatajczyk/Langmal-Server/models"

type QuizRepoInterface interface {
	FindAll() []models.QuizEntity
	Find(id string) (models.QuizEntity, bool)
}

func NewQuizRepo() QuizRepoInterface {
	quizzes := []models.QuizEntity{createQuiz1(), createQuiz2()}
	return &quizRepo{
		quizzes: quizzes,
	}
}

type quizRepo struct {
	quizzes []models.QuizEntity
}

func (qr *quizRepo) FindAll() []models.QuizEntity {
	return qr.quizzes
}

func (qr *quizRepo) Find(id string) (models.QuizEntity, bool) {
	for _, quiz := range qr.quizzes {
		if quiz.Id == id {
			return quiz, true
		}
	}
	return models.QuizEntity{}, false
}

// TODO: remove and use DB instead
func createQuiz1() models.QuizEntity {
	question1 := models.QuestionEntity{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer A",
	}
	question2 := models.QuestionEntity{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}
	question3 := models.QuestionEntity{
		Title:   "Third question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}

	quiz := models.QuizEntity{
		Title:     "First quiz",
		Id:        "4e2778d3-57df-4fe9-83ec-af5ffec1ec5c",
		Questions: []models.QuestionEntity{question1, question2, question3},
	}

	return quiz
}

// TODO: remove and use DB instead
func createQuiz2() models.QuizEntity {
	question1 := models.QuestionEntity{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}
	question2 := models.QuestionEntity{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}

	quiz := models.QuizEntity{
		Title:     "Second quiz",
		Id:        "5e8ef788-f305-4ee3-ad69-ba8924ca3806",
		Questions: []models.QuestionEntity{question1, question2},
	}

	return quiz
}
