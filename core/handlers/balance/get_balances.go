package balanceHandler

import (
	"net/http"
	"tracker/core/logger"

	"github.com/gin-gonic/gin"
)

func (c *BalanceHandler) GetBalances(ctx *gin.Context) {

	b, err := c.DB.Balance.Query().All(ctx)
	if err != nil {
		logger.Err("failed getting balances: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't get balances")
	}
	logger.Log("completed balances search")

	ctx.JSON(http.StatusOK, b)
}
