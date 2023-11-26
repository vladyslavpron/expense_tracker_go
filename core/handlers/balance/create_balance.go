package balanceHandler

import (
	"net/http"
	"tracker/core/logger"

	"github.com/gin-gonic/gin"
)

func (c *BalanceHandler) CreateBalance(ctx *gin.Context) {
	var dto CreateBalanceDTO
	if err := ctx.BindJSON(&dto); err != nil {
		logger.Err("ERR: CreateBalance func" + string(err.Error()))
		ctx.JSON(http.StatusBadRequest, "Wrong DTO")
		return
	}

	b, err := c.DB.Balance.Create().SetTitle(dto.Title).Save(ctx)
	if err != nil {
		logger.Err("failed creating balance: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't create balance")

	}
	logger.Log("balance was created: " + b.String())

	ctx.JSON(http.StatusCreated, b)
}

type CreateBalanceDTO struct {
	Title string
}
