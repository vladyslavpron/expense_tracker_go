package transactionHandler

import (
	"fmt"
	"net/http"
	"strconv"
	"tracker/core/logger"
	"tracker/ent/transaction"

	"github.com/gin-gonic/gin"
)

func (c *TransactionHandler) GetTransactions(ctx *gin.Context) {

	balanceId, bIdPresent := ctx.GetQuery("balance_id")
	categoryId, cIdPresent := ctx.GetQuery("category_id")
	// params to get transactions per category or balance, with pagination?

	// need to return count if with pagination
	query := c.DB.Transaction.Query()
	if bId, err := strconv.Atoi(balanceId); err == nil && bIdPresent {
		query.Where(transaction.BalanceID(bId))
	}
	if cId, err := strconv.Atoi(categoryId); err == nil && cIdPresent {
		query.Where(transaction.CategoryID(cId))
	}
	t, err := query.All(ctx)
	if err != nil {
		logger.Err("failed getting transactions: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't get transactions")
	}
	logger.Log("completed transactions search")
	fmt.Println(t)

	ctx.JSON(http.StatusOK, t)
}
