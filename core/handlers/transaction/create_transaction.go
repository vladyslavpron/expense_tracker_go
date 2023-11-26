package transactionHandler

import (
	"net/http"
	"tracker/core/logger"

	"github.com/gin-gonic/gin"
)

func (c *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var dto CreateTransactionDTO
	if err := ctx.BindJSON(&dto); err != nil {
		logger.Err("ERR: CreateTransaction func" + string(err.Error()))
		ctx.JSON(http.StatusBadRequest, "Wrong DTO")
		return
	}

	t, err := c.DB.Transaction.Create().SetDescription(dto.Description).SetAmount(dto.Amount).SetBalanceID(dto.BalanceId).SetCategoryID(dto.CategoryId).Save(ctx)
	if err != nil {
		logger.Err("failed creating transaction: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't create transaction")
	}
	logger.Log("transaction was created: " + t.String())

	ctx.JSON(http.StatusCreated, t)
}

type CreateTransactionDTO struct {
	Description string
	Amount      float32
	CategoryId  int
	BalanceId   int
}
