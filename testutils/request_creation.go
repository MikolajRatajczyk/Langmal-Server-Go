package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePostJsonRequest(jsonMap gin.H) *http.Request {
	readCloser := createReadCloser(jsonMap)
	if readCloser == nil {
		return nil
	}

	header := http.Header{}
	header.Set("Content-Type", "application/json")

	request := http.Request{
		Body:   readCloser,
		Method: http.MethodPost,
		Header: header,
	}

	return &request
}

func CreateEmptyGetRequest() *http.Request {
	request, err := http.NewRequest(http.MethodGet, "/foo", nil)
	if err != nil {
		return nil
	}

	return request
}

func createReadCloser(jsonMap gin.H) io.ReadCloser {
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		return nil
	}

	jsonBuffer := bytes.NewBuffer(jsonBytes)
	readCloser := io.NopCloser(jsonBuffer)
	return readCloser
}
