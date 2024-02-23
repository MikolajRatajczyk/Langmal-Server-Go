package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/testutils"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

var validResults = gin.H{"correct": 1, "wrong": 1, "quiz_id": "foo", "created_at": 1}

func TestResultsController_SaveResultsWhenRequestOk(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: true}

	testResultsController_SaveResults(&service, &extractor, validResults, true, 201, t)
}

func TestResultsController_SaveResultsWhenRequestInvalid(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: true}
	invalidResults := gin.H{"correct": 1}

	testResultsController_SaveResults(&service, &extractor, invalidResults, true, 400, t)
}

func TestResultsController_SaveResultsWhenRequestBodyEmpty(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: true}

	testResultsController_SaveResults(&service, &extractor, gin.H{}, true, 400, t)
}

func TestResultsController_SaveResultsWhenNoAuthHeader(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: true}

	testResultsController_SaveResults(&service, &extractor, validResults, false, 400, t)
}

func TestResultsController_SaveResultsWhenNoUserIdInClaims(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: false}

	testResultsController_SaveResults(&service, &extractor, validResults, true, 401, t)
}

func TestResultsController_SaveResultsWhenServiceFails(t *testing.T) {
	service := resultServiceFake{saveSuccessful: false}
	extractor := claimsExtractorFake{successful: true}

	testResultsController_SaveResults(&service, &extractor, validResults, true, 500, t)
}

func testResultsController_SaveResults(
	service services.ResultServiceInterface,
	extractor utils.ClaimsExtractorInterface,
	requestBody gin.H,
	includeAuthHeader bool,
	expectedCode int,
	t *testing.T,
) {
	request := testutils.CreatePostJsonRequest(requestBody)
	if request == nil {
		t.Fatal("Can't create a request and continue the test")
	}
	if includeAuthHeader {
		request.Header.Set("Authorization", "Bearer foo")
	}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	sut := ResultsController{
		ResultService:   service,
		ClaimsExtractor: extractor,
	}

	sut.SaveResult(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}
}

type resultServiceFake struct {
	saveSuccessful bool
}

func (rsf *resultServiceFake) Save(result models.ResultWriteDto, userId string) bool {
	return rsf.saveSuccessful
}

func (*resultServiceFake) Find(userId string) []models.ResultReadDto {
	result := models.ResultReadDto{
		ResultWriteDto: models.ResultWriteDto{
			Correct:   1,
			Wrong:     1,
			QuizId:    "foo",
			CreatedAt: 1,
		},
		QuizTitle: "bar",
	}

	return []models.ResultReadDto{result}
}
