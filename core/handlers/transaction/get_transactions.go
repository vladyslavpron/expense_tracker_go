package transactionHandler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"tracker/core/logger"
	"tracker/ent"
	"tracker/ent/transaction"

	"github.com/gin-gonic/gin"
)

func (c *TransactionHandler) GetTransactionsHandler(ctx *gin.Context) {

	balanceIdString, _ := ctx.GetQuery("balance_id")
	categoryIdString, _ := ctx.GetQuery("category_id")
	balanceId, _ := strconv.Atoi(balanceIdString)
	categoryId, _ := strconv.Atoi(categoryIdString)

	t, err := c.GetTransactions(ctx, GetTransactionsParams{BalanceID: balanceId, CategoryId: categoryId})
	if err != nil {
		logger.Err("failed getting transactions: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't get transactions")
	}
	logger.Log("completed transactions search")
	fmt.Println(t)

	ctx.JSON(http.StatusOK, t)
}

func (c *TransactionHandler) GetTransactions(ctx context.Context, params GetTransactionsParams) ([]*ent.Transaction, error) {

	// need to return count if with pagination
	query := c.DB.Transaction.Query()
	if params.BalanceID != 0 {
		query.Where(transaction.BalanceID(params.BalanceID))
	}
	if params.CategoryId != 0 {
		query.Where(transaction.CategoryID(params.CategoryId))
	}
	t, err := query.All(ctx)
	if err != nil {
		logger.Err("failed getting transactions: " + err.Error())
		return nil, err
	}
	logger.Log("completed transactions search")

	return t, nil
}

type GetTransactionsParams struct {
	BalanceID  int
	CategoryId int
}
