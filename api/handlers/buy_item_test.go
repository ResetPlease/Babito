package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/ResetPlease/Babito/internal/models"
	testcore "github.com/ResetPlease/Babito/internal/test_core"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestBuyItem(t *testing.T) {
	router := gin.Default()
	testCore := testcore.NewTestCore()
	err := testCore.DB.TestClearOperationHistory()
	assert.Equal(t, err, nil)
	// for create users
	router.POST("/api/auth", testCore.Handler.AuthHandler)
	secure := router.Group("/api").Use(testCore.Middleware.AuthMiddleware())
	{
		secure.GET("/buy/:item", testCore.Handler.BuyItemHandler)
	}

	user := models.AuthRequest{
		Username: "user",
		Password: "pass",
	}

	token, err := testCore.CreateTestUserWithToken(user, router)
	assert.Equal(t, err, nil)

	t.Run("buy_item_ok", func(t *testing.T) {
		err := testCore.DB.TestUpdateUsersBalance()
		assert.Equal(t, err, nil)

		product := models.BOOK
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/buy/%s", product), nil)
		assert.Equal(t, err, nil)
		testCore.SetAuthToken(request, *token.Token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		assert.Equal(t, recorder.Code, http.StatusOK)

		userData, err := testCore.DB.GetUserDataByUsername(user.Username)
		assert.Equal(t, err, nil)
		assert.Equal(t, userData.Balance, testCore.Middleware.Config.DefaultUserBalance-product.Price())

		operations, err := testCore.DB.GetPurchaseByUserID(userData.ID)
		assert.Equal(t, err, nil)
		assert.Equal(t, len(operations), 1)
		assert.Equal(t, operations[0].Amount, product.Price())
		assert.Equal(t, operations[0].Item, string(product))
	})

	t.Run("test_error_responses", func(t *testing.T) {
		testCases := []struct {
			Name            string
			Item            string
			Times           int
			ExpectedMessage string
		}{
			{
				Name:            "test_product_not_found",
				Item:            string(models.FAKEPRODUCT),
				Times:           1,
				ExpectedMessage: "Product not found",
			},
			{
				Name:            "test_not_enought_funds",
				Item:            string(models.PINKHOODY),
				Times:           3,
				ExpectedMessage: "Not enought funds",
			},
		}

		for _, tt := range testCases {
			t.Run(tt.Name, func(t *testing.T) {
				err := testCore.DB.TestUpdateUsersBalance()
				assert.Equal(t, err, nil)

				url := fmt.Sprintf("/api/buy/%s", tt.Item)
				request, err := http.NewRequest(http.MethodGet, url, nil)
				testCore.SetAuthToken(request, *token.Token)
				assert.Equal(t, err, nil)

				var recorder *httptest.ResponseRecorder
				for i := 0; i < tt.Times; i++ {
					recorder = httptest.NewRecorder()
					router.ServeHTTP(recorder, request)
				}

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
