package models

type QuizEntity struct {
	Title string
	// must match QuestionEntity.QuizEntityId
	Id        string           `gorm:"primaryKey"`
	Questions []QuestionEntity `gorm:"foreignKey:QuizEntityId"`
}

func GiantsQuiz() QuizEntity {
	const quizId = "4e2778d3-57df-4fe9-83ec-afffec1ec5c"

	mountain := QuestionEntity{
		Title:        "What is the highest mountain 🏔️?",
		Id:           "1953be72-bd01-4a94-916e-e7d5afb673c2",
		Options:      []string{"Nanga Parbat", "Mount Everest", "Galdhøpiggen"},
		Answer:       1,
		QuizEntityId: quizId,
	}
	city := QuestionEntity{
		Title:        "What is the most populated city 🌇?",
		Id:           "3c472c83-c935-4b2a-8840-150c7f7e1162",
		Options:      []string{"Tokyo", "London", "São Paulo"},
		Answer:       0,
		QuizEntityId: quizId,
	}
	animal := QuestionEntity{
		Title:        "What is the heaviest animal ⚖️?",
		Id:           "f35504fa-0544-4976-be61-08b27df50a5a",
		Options:      []string{"Great White Shark", "African Elephant", "Blue Whale"},
		Answer:       2,
		QuizEntityId: quizId,
	}

	quiz := QuizEntity{
		Title:     "Giants of the world 🌍",
		Id:        quizId,
		Questions: []QuestionEntity{mountain, city, animal},
	}

	return quiz
}

func SpaceQuiz() QuizEntity {
	const quizId = "5e8ef788-f305-4ee3-ad69-ba8924ca3806"

	redPlanet := QuestionEntity{
		Title:        `Which planet is known as the "Red Planet" 🔴?`,
		Id:           "3ba01efa-e7e7-4d3e-9cb3-b3f21016e7b7",
		Options:      []string{"Jupiter", "Saturn", "Mars"},
		Answer:       2,
		QuizEntityId: quizId,
	}
	galaxy := QuestionEntity{
		Title:        "What galaxy is Earth located in 🌌?",
		Id:           "6a4d0e7f-ea00-49ef-8aea-9bb3a9041a7e",
		Options:      []string{"Andromeda", "Milky Way", "Large Magellanic Cloud"},
		Answer:       1,
		QuizEntityId: quizId,
	}

	quiz := QuizEntity{
		Title:     "Space and beyond 🚀",
		Id:        quizId,
		Questions: []QuestionEntity{redPlanet, galaxy},
	}

	return quiz
}
