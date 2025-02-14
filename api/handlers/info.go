package handlers

import (
	"log/slog"
	"net/http"

	models "github.com/ResetPlease/Babito/internal/models"
	"github.com/ResetPlease/Babito/internal/tools"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InfoHanlder(c *gin.Context) {
	user, err := tools.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorUnauthorized)
		return
	}
	operations, err := h.db.GetAllUserOperations(user.ID)
	if err != nil {
		h.logger.Error("Got error while get user info", slog.Any("user", user), slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, models.ErrorInternalServerError)
		return
	}

	userData, err := h.db.GetUserDataByUserID(user.ID)
	if err != nil {
		h.logger.Error("Got error while get user data(balance)", slog.Any("user", user), slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, models.ErrorInternalServerError)
		return
	}

	var response models.InfoResponse
	intUserBalance := int(userData.Balance)
	response.Coins = &intUserBalance

	items := make(map[string]int)

	for _, operation := range operations {
		// current user buy item
		if operation.Type == models.PURCHASE {
			items[operation.Item.String]++
			continue
		}

		// if current user(taget_user_id) receive coins from user(user_id)
		if operation.Type == models.TRANSFER && operation.TargetUserID.Int64 == int64(user.ID) {
			if response.CoinHistory == nil {
				response.CoinHistory = new(models.CoinHistoryField)
			}

			if response.CoinHistory.Received == nil {
				receivedObj := make(models.RecivedField, 0)
				response.CoinHistory.Received = &receivedObj
			}

			(*response.CoinHistory.Received) = append(
				(*response.CoinHistory.Received),
				models.RecivedItem{
					Amount:   int(operation.Amount),
					FromUser: operation.Username,
				},
			)
			continue
		}

		// if current user transfer coin to target user
		if operation.Type == models.TRANSFER && operation.UserID == user.ID {
			if response.CoinHistory == nil {
				response.CoinHistory = new(models.CoinHistoryField)
			}

			if response.CoinHistory.Sent == nil {
				sentObj := make(models.SentField, 0)
				response.CoinHistory.Sent = &sentObj
			}

			(*response.CoinHistory.Sent) = append(
				(*response.CoinHistory.Sent),
				models.SentItem{
					Amount: int(operation.Amount),
					ToUser: operation.TargetUsername.String,
				},
			)
		}
	}

	for name, quantity := range items {
		if response.Inventory == nil {
			inventoryObj := make(models.InventoryField, 0)
			response.Inventory = &inventoryObj
		}

		(*response.Inventory) = append(
			(*response.Inventory),
			models.InventoryItem{
				Quantity: quantity,
				Type:     name,
			},
		)
	}

	c.JSON(http.StatusOK, response)
}
