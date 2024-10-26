package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/models"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/services"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/testutils"
	"github.com/gin-gonic/gin"
)

func TestQuizController_GetQuizzesWhenServiceOk(t *testing.T) {
	service := fakeQuizService{successful: true}
	testQuizController_GetQuizzes(&service, 200, true, t)
}

func TestQuizController_GetQuizzesWhenServiceFails(t *testing.T) {
	service := fakeQuizService{successful: false}
	testQuizController_GetQuizzes(&service, 404, false, t)
}

func testQuizController_GetQuizzes(
	service services.QuizServiceInterface,
	expectedCode int,
	notEmptyParsableResponseExpected bool,
	t *testing.T,
) {
	request := testutils.CreateEmptyGetRequest()
	if request == nil {
		t.Fatal("Can't create a request and continue the test")
	}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	sut := QuizController{Service: service}

	sut.GetQuizzes(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}

	if notEmptyParsableResponseExpected {
		dtos, ok := parseQuizDtos(recorder.Body)
		if !ok {
			t.Errorf("Can't create DTO array from the response body")
		}
		if len(dtos) == 0 {
			t.Errorf("DTO array from the response body should not be empty")
		}
	}
}

func parseQuizDtos(body *bytes.Buffer) ([]models.QuizDto, bool) {
	var parsed []models.QuizDto
	err := json.Unmarshal(body.Bytes(), &parsed)
	if err != nil {
		return []models.QuizDto{}, false
	}

	return parsed, true
}

type fakeQuizService struct {
	successful bool
}

func (fqs *fakeQuizService) All() ([]models.QuizDto, bool) {
	if !fqs.successful {
		return []models.QuizDto{}, false
	}

	question := models.QuestionDto{
		Title:   "foo",
		Options: []string{"bar"},
		Answer:  0,
	}

	quiz := models.QuizDto{
		Title:     "baz",
		Id:        "abc",
		Questions: []models.QuestionDto{question},
	}

	return []models.QuizDto{quiz}, true
}
