package repositories

import "github.com/MikolajRatajczyk/Langmal-Server/models"

type QuizRepoInterface interface {
	FindAll() []models.Quiz
}

func NewQuizRepo() QuizRepoInterface {
	//	TODO: open DB session and pass it to quizRepo
	return &quizRepo{}
}

type quizRepo struct {
	//	TODO: add `connection *gorm.DB` or similar and use it
}

func (qr *quizRepo) FindAll() []models.Quiz {
	quizzes := []models.Quiz{createQuiz1(), createQuiz2()}
	return quizzes
}

// TODO: remove and use DB instead
func createQuiz1() models.Quiz {
	question1 := models.Question{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer A",
	}
	question2 := models.Question{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}
	question3 := models.Question{
		Title:   "Third question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}

	quiz := models.Quiz{
		Title:     "First quiz",
		Id:        "4e2778d3-57df-4fe9-83ec-af5ffec1ec5c",
		Questions: []models.Question{question1, question2, question3},
	}

	return quiz
}

func createQuiz2() models.Quiz {
	question1 := models.Question{
		Title:   "First question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer C",
	}
	question2 := models.Question{
		Title:   "Second question from the server",
		Options: []string{"Answer A", "Answer B", "Answer C"},
		Answer:  "Answer B",
	}

	quiz := models.Quiz{
		Title:     "Second quiz",
		Id:        "5e8ef788-f305-4ee3-ad69-ba8924ca3806",
		Questions: []models.Question{question1, question2},
	}

	return quiz
}
