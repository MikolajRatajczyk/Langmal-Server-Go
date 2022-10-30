package repositories

import "github.com/MikolajRatajczyk/Langmal-Server/entities"

type TestRepository interface {
	FindAll() []entities.Test
}

func NewTestRepository() TestRepository {
	//	TODO: open DB session and pass it to testRepository
	return &testRepository{}
}

type testRepository struct {
	//	TODO: add `connection *gorm.DB` or similar and use it
}

func (tr *testRepository) FindAll() []entities.Test {
	test := createTest()
	return []entities.Test{test}
}

// TODO: remove and use DB instead
func createTest() entities.Test {
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

	test := entities.Test{
		Name:      "First test",
		Id:        "4e2778d3-57df-4fe9-83ec-af5ffec1ec5c",
		Questions: []entities.Question{question1, question2, question3},
	}

	return test
}
