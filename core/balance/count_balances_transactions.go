package balance

import (
	"context"
	"fmt"
	"tracker/core/logger"
	"tracker/ent"
	"tracker/ent/transaction"
)

func CountBalancesTransactions(ctx context.Context, db *ent.Client) ([]TransactionSumCount, error) {

	var r []TransactionSumCount

	err := db.Transaction.Query().GroupBy(transaction.FieldBalanceID).Aggregate(
		ent.Count(),
		ent.Sum(transaction.FieldAmount),
	).Scan(ctx, &r)

	if err != nil {
		logger.Err("Error when calculating count of transactions for balance")
		return nil, err
	}

	fmt.Println(r)
	return r, nil
}

type TransactionSumCount struct {
	BalanceID int     `json:"balance_id"`
	Sum       float64 `json:"sum"`
	Count     int     `json:"count"`
}
