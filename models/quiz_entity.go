package models

type QuizEntity struct {
	Title string
	// must match QuestionEntity.QuizEntityId
	Id        string           `gorm:"primaryKey"`
	Questions []QuestionEntity `gorm:"foreignKey:QuizEntityId"`
}

func DefaultQuiz1() QuizEntity {
	const quizId = "4e2778d3-57df-4fe9-83ec-afffec1ec5c"

	question1 := QuestionEntity{
		Title:        "First question from the server",
		Id:           "1953be72-bd01-4a94-916e-e7d5afb673c2",
		Options:      []string{"Answer A", "Answer B", "Answer C"},
		Answer:       0,
		QuizEntityId: quizId,
	}
	question2 := QuestionEntity{
		Title:        "Second question from the server",
		Id:           "3c472c83-c935-4b2a-8840-150c7f7e1162",
		Options:      []string{"Answer A", "Answer B", "Answer C"},
		Answer:       1,
		QuizEntityId: quizId,
	}
	question3 := QuestionEntity{
		Title:        "Third question from the server",
		Id:           "f35504fa-0544-4976-be61-08b27df50a5a",
		Options:      []string{"Answer A", "Answer B", "Answer C"},
		Answer:       2,
		QuizEntityId: quizId,
	}

	quiz := QuizEntity{
		Title:     "First default quiz",
		Id:        quizId,
		Questions: []QuestionEntity{question1, question2, question3},
	}

	return quiz
}

func DefaultQuiz2() QuizEntity {
	const quizId = "5e8ef788-f305-4ee3-ad69-ba8924ca3806"

	question1 := QuestionEntity{
		Title:        "First question from the server",
		Id:           "3ba01efa-e7e7-4d3e-9cb3-b3f21016e7b7",
		Options:      []string{"Answer A", "Answer B", "Answer C"},
		Answer:       2,
		QuizEntityId: quizId,
	}
	question2 := QuestionEntity{
		Title:        "Second question from the server",
		Id:           "6a4d0e7f-ea00-49ef-8aea-9bb3a9041a7e",
		Options:      []string{"Answer A", "Answer B", "Answer C"},
		Answer:       1,
		QuizEntityId: quizId,
	}

	quiz := QuizEntity{
		Title:     "Second default quiz",
		Id:        quizId,
		Questions: []QuestionEntity{question1, question2},
	}

	return quiz
}
