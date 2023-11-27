package balanceHandler

import (
	"context"
	"net/http"
	"tracker/core/logger"
	"tracker/ent"

	"github.com/gin-gonic/gin"
)

func (c *BalanceHandler) GetBalancesHandler(ctx *gin.Context) {

	b, err := c.GetBalances(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Can't get balances")
		return
	}

	ctx.JSON(http.StatusOK, b)
}

func (c *BalanceHandler) GetBalances(ctx context.Context) ([]*ent.Balance, error) {
	b, err := c.DB.Balance.Query().All(ctx)
	if err != nil {
		logger.Err("failed getting balances: " + err.Error())
		return nil, err
	}
	logger.Log("completed balances search")
	return b, nil
}
