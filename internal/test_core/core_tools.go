package testcore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/ResetPlease/Babito/internal/models"
	"github.com/gin-gonic/gin"
)

func (tc *TestCore) CreateTestUserWithToken(userCreds models.AuthRequest, router *gin.Engine) (*models.AuthResponse, error) {
	body, err := json.Marshal(userCreds)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "/api/auth", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", recorder.Code)
	}

	if recorder.Body == nil {
		return nil, errors.New("empty response body")
	}

	rawData, err := io.ReadAll(recorder.Body)
	if err != nil {
		return nil, err
	}

	var result models.AuthResponse
	err = json.Unmarshal(rawData, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (tc *TestCore) SetAuthToken(request *http.Request, token string) {
	request.Header.Set(models.AuthHeaderKey, fmt.Sprintf("Bearer %s", token))
}
