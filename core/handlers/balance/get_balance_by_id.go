package balanceHandler

import (
	"context"
	"tracker/core/logger"
	"tracker/ent"
	balanceEntity "tracker/ent/balance"
)

func (c *BalanceHandler) GetBalanceById(ctx context.Context, id int) (*ent.Balance, error) {
	b, err := c.DB.Balance.Query().Where(balanceEntity.ID(id)).First(ctx)
	if err != nil {
		logger.Err("failed getting balance by id: " + err.Error())
		return nil, err
	}
	logger.Log("completed balance by id search")
	return b, nil
}
