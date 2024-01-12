package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/testutils"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

func TestResultsController_GetResultsWhenRequestOk(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: true}

	testResultsController_GetResults(&service, &extractor, true, 200, true, t)
}

func TestResultsController_GetResultsWhenNoAuthHeader(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: true}

	testResultsController_GetResults(&service, &extractor, false, 400, false, t)
}

func TestResultsController_GetResultsWhenNoUserIdInClaims(t *testing.T) {
	service := resultServiceFake{saveSuccessful: true}
	extractor := claimsExtractorFake{successful: false}

	testResultsController_GetResults(&service, &extractor, true, 401, false, t)
}

func testResultsController_GetResults(
	service services.ResultServiceInterface,
	extractor utils.ClaimsExtractorInterface,
	includeAuthHeader bool,
	expectedCode int,
	notEmptyParsableResponseExpected bool,
	t *testing.T,
) {
	request := testutils.CreateEmptyGetRequest()
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

	sut.GetResults(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}

	if notEmptyParsableResponseExpected {
		dtos, ok := parseResultDtos(recorder.Body)
		if !ok {
			t.Errorf("Can't create DTO array from the given body")
		}
		if len(dtos) == 0 {
			t.Errorf("DTO array should not be empty")
		}
	}
}

func parseResultDtos(body *bytes.Buffer) ([]models.ResultDto, bool) {
	var parsed []models.ResultDto
	err := json.Unmarshal(body.Bytes(), &parsed)
	if err != nil {
		return []models.ResultDto{}, false
	}

	return parsed, true
}
