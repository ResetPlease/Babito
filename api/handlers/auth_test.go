package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ResetPlease/Babito/internal/models"
	testcore "github.com/ResetPlease/Babito/internal/test_core"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestAuthHandler(t *testing.T) {
	router := gin.Default()
	testCore := testcore.NewTestCore()
	router.POST("/api/auth", testCore.Handler.AuthHandler)

	t.Run("test_status_ok", func(t *testing.T) {
		userData := models.AuthRequest{
			Username: "test",
			Password: "test",
		}

		data, err := json.Marshal(userData)
		assert.Equal(t, err, nil)

		req, err := http.NewRequest(http.MethodPost, "/api/auth", strings.NewReader(string(data)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var responseAuthData models.AuthResponse
		err = json.Unmarshal(rr.Body.Bytes(), &responseAuthData)
		assert.Equal(t, err, nil)

		JWTRegexp := `^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+$`
		assert.MatchRegex(t, *responseAuthData.Token, JWTRegexp)
	})

	t.Run("test_bad_request_wrong_data_format", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/auth", strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{"errors":"Wrong data format"}`
		assert.Equal(t, expected, rr.Body.String())
	})

	t.Run("test_bad_request_missing_required_field", func(t *testing.T) {
		userData := models.AuthRequest{
			Username: "test",
			Password: "",
		}

		data, err := json.Marshal(userData)
		assert.Equal(t, err, nil)

		req, err := http.NewRequest(http.MethodPost, "/api/auth", strings.NewReader(string(data)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{"errors":"Missing required field"}`
		assert.Equal(t, expected, rr.Body.String())
	})

	t.Run("test_unauthorized_wrong_password", func(t *testing.T) {
		userData := models.AuthRequest{
			Username: "test", // already exist
			Password: "wrong_password",
		}

		data, err := json.Marshal(userData)
		assert.Equal(t, err, nil)

		req, err := http.NewRequest(http.MethodPost, "/api/auth", strings.NewReader(string(data)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		expected := `{"errors":"Unauthorized"}`
		assert.Equal(t, expected, rr.Body.String())
	})
}
