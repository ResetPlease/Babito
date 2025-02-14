package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ResetPlease/Babito/internal/models"
	testcore "github.com/ResetPlease/Babito/internal/test_core"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestSendCoin(t *testing.T) {
	router := gin.Default()
	testCore := testcore.NewTestCore()
	err := testCore.DB.TestClearOperationHistory()
	assert.Equal(t, err, nil)
	// for create users
	router.POST("/api/auth", testCore.Handler.AuthHandler)
	secure := router.Group("/api").Use(testCore.Middleware.AuthMiddleware())
	{
		secure.POST("/sendCoin", testCore.Handler.SendCoinHandler)
	}

	firstUser := models.AuthRequest{
		Username: "first",
		Password: "pass",
	}

	secondUser := models.AuthRequest{
		Username: "second",
		Password: "pass",
	}

	firstToken, err := testCore.CreateTestUserWithToken(firstUser, router)
	assert.Equal(t, err, nil)
	_, err = testCore.CreateTestUserWithToken(secondUser, router)
	assert.Equal(t, err, nil)

	user, err := tools.ParseJWTToken(*firstToken.Token, testCore.Middleware.Config)
	assert.Equal(t, err, nil)

	t.Run("test_send_ok", func(t *testing.T) {
		err := testCore.DB.TestUpdateUsersBalance()
		assert.Equal(t, err, nil)
		payload := models.SendCoinRequest{
			Amount: 100,
			ToUser: secondUser.Username,
		}
		jsonPayload, err := json.Marshal(payload)
		assert.Equal(t, err, nil)

		request, err := http.NewRequest(http.MethodPost, "/api/sendCoin", strings.NewReader(string(jsonPayload)))
		testCore.SetAuthToken(request, *firstToken.Token)
		assert.Equal(t, err, nil)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		operations, err := testCore.DB.GetTransfersByUserID(user.ID)
		assert.Equal(t, err, nil)
		assert.Equal(t, len(operations), 1)
		assert.Equal(t, operations[0].Amount, int64(100))
		assert.Equal(t, operations[0].TargetUsername.String, secondUser.Username)
	})

	t.Run("test_error_responses", func(t *testing.T) {
		testCases := []struct {
			Name            string
			Payload         models.SendCoinRequest
			ExpectedMessage string
		}{
			{
				Name: "test_negative_amount",
				Payload: models.SendCoinRequest{
					Amount: -200,
					ToUser: secondUser.Username,
				},
				ExpectedMessage: "Amount should be more than zero",
			},
			{
				Name: "test_selftransfer",
				Payload: models.SendCoinRequest{
					Amount: 100,
					ToUser: firstUser.Username,
				},
				ExpectedMessage: "Replenishment of your account is prohibited",
			},
			{
				Name: "test_not_enough_funds",
				Payload: models.SendCoinRequest{
					Amount: 2000,
					ToUser: secondUser.Username,
				},
				ExpectedMessage: "Not enought funds",
			},
		}

		for _, tt := range testCases {
			t.Run(tt.Name, func(t *testing.T) {
				err := testCore.DB.TestUpdateUsersBalance()
				assert.Equal(t, err, nil)
				jsonPayload, err := json.Marshal(tt.Payload)
				assert.Equal(t, err, nil)

				request, err := http.NewRequest(http.MethodPost, "/api/sendCoin", strings.NewReader(string(jsonPayload)))
				testCore.SetAuthToken(request, *firstToken.Token)
				assert.Equal(t, err, nil)

				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, request)

				assert.Equal(t, http.StatusBadRequest, recorder.Code)

				var response models.ErrorResponse
				assert.NotEqual(t, recorder.Body, nil)
				err = json.Unmarshal(recorder.Body.Bytes(), &response)
				assert.Equal(t, err, nil)
				assert.Equal(t, response.Errors, tt.ExpectedMessage)
			})
		}
	})

}
