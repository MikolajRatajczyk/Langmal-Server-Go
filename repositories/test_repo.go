package repositories

import "github.com/MikolajRatajczyk/Langmal-Server/models"

type TestRepoInterface interface {
	FindAll() []models.Test
}

func NewTestRepo() TestRepoInterface {
	//	TODO: open DB session and pass it to testRepo
	return &testRepo{}
}

type testRepo struct {
	//	TODO: add `connection *gorm.DB` or similar and use it
}

func (tr *testRepo) FindAll() []models.Test {
	tests := []models.Test{createTest1(), createTest2()}
	return tests
}

// TODO: remove and use DB instead
func createTest1() models.Test {
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

	test := models.Test{
		Name:      "First test",
		Id:        "4e2778d3-57df-4fe9-83ec-af5ffec1ec5c",
		Questions: []models.Question{question1, question2, question3},
	}

	return test
}

func createTest2() models.Test {
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

	test := models.Test{
		Name:      "Second test",
		Id:        "5e8ef788-f305-4ee3-ad69-ba8924ca3806",
		Questions: []models.Question{question1, question2},
	}

	return test
}
