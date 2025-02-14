package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/ResetPlease/Babito/internal/models"
	testcore "github.com/ResetPlease/Babito/internal/test_core"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestInfoHanlder(t *testing.T) {
	router := gin.Default()
	testCore := testcore.NewTestCore()

	// for create users
	router.POST("/api/auth", testCore.Handler.AuthHandler)
	secure := router.Group("/api").Use(testCore.Middleware.AuthMiddleware())
	{
		secure.POST("/sendCoin", testCore.Handler.SendCoinHandler)
		secure.POST("/buy/:item", testCore.Handler.BuyItemHandler)
		secure.GET("/info", testCore.Handler.InfoHanlder)
	}

	UserA := models.AuthRequest{
		Username: "first",
		Password: "pass",
	}

	UserB := models.AuthRequest{
		Username: "second",
		Password: "pass",
	}

	tokenA, err := testCore.CreateTestUserWithToken(UserA, router)
	assert.Equal(t, err, nil)
	tokenB, err := testCore.CreateTestUserWithToken(UserB, router)
	assert.Equal(t, err, nil)

	fromUser, err := tools.ParseJWTToken(*tokenA.Token, testCore.Middleware.Config)
	assert.Equal(t, err, nil)
	toUser, err := tools.ParseJWTToken(*tokenB.Token, testCore.Middleware.Config)
	assert.Equal(t, err, nil)

	t.Run("test_info_ok", func(t *testing.T) {
		err := testCore.DB.TestClearOperationHistory()
		assert.Equal(t, err, nil)
		err = testCore.DB.TestUpdateUsersBalance()
		assert.Equal(t, err, nil)

		request, err := http.NewRequest(http.MethodGet, "/api/info", nil)
		testCore.SetAuthToken(request, *tokenA.Token)
		assert.Equal(t, err, nil)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		var response models.InfoResponse
		assert.NotEqual(t, recorder.Body, nil)
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.Equal(t, err, nil)

		assert.Equal(t, *response.Coins, int(testCore.Middleware.Config.DefaultUserBalance))
		assert.Equal(t, response.CoinHistory, nil)
	})

	t.Run("test_with_transfer_and_receive", func(t *testing.T) {
		err := testCore.DB.TestClearOperationHistory()
		assert.Equal(t, err, nil)
		err = testCore.DB.TestUpdateUsersBalance()
		assert.Equal(t, err, nil)

		amount := 500
		err = testCore.DB.SendCoinByUsername(fromUser.ID, toUser.Username, int64(amount))
		assert.Equal(t, err, nil)

		testCases := []struct {
			Name            string
			Token           string
			ExpectedBalance int64
		}{
			{
				Name:            "sent",
				Token:           *tokenA.Token,
				ExpectedBalance: int64(amount),
			},
			{
				Name:            "received",
				Token:           *tokenB.Token,
				ExpectedBalance: testCore.Middleware.Config.DefaultUserBalance + int64(amount),
			},
		}

		for _, tt := range testCases {
			t.Run(tt.Name, func(t *testing.T) {
				request, err := http.NewRequest(http.MethodGet, "/api/info", nil)
				testCore.SetAuthToken(request, tt.Token)
				assert.Equal(t, err, nil)

				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, request)
				assert.Equal(t, recorder.Code, http.StatusOK)

				var response models.InfoResponse
				assert.NotEqual(t, recorder.Body, nil)
				err = json.Unmarshal(recorder.Body.Bytes(), &response)
				assert.Equal(t, err, nil)

				assert.Equal(t, *response.Coins, int(tt.ExpectedBalance))

				if tt.Name == "sent" {
					// check sent coins
					assert.Equal(t, len(*response.CoinHistory.Sent), 1)
					assert.Equal(t, (*response.CoinHistory.Sent)[0].Amount, amount)
					assert.Equal(t, (*response.CoinHistory.Sent)[0].ToUser, toUser.Username)
				} else if tt.Name == "received" {
					// check received coins
					assert.Equal(t, len(*response.CoinHistory.Received), 1)
					assert.Equal(t, (*response.CoinHistory.Received)[0].Amount, amount)
					assert.Equal(t, (*response.CoinHistory.Received)[0].FromUser, fromUser.Username)
				}
			})
		}
	})

}
