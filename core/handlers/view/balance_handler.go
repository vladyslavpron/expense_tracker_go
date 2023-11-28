package viewHandler

import (
	"strconv"
	transactionHandler "tracker/core/handlers/transaction"
	"tracker/ent"

	"github.com/gin-gonic/gin"
)

func (h *ViewHandler) Balance(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	balance, _ := h.BalanceHandler.GetBalanceById(ctx, id)

	transactions, _ := h.TransactionHandler.GetTransactions(ctx, transactionHandler.GetTransactionsParams{BalanceID: id})

	categories, _ := h.CategoryHandler.GetCategories(ctx)
	categoryTitles := make(map[int]string, len(categories))
	for _, category := range categories {
		categoryTitles[category.ID] = category.Title
	}

	positiveSum := 0.0
	positiveSumCount := 0
	negativeSum := 0.0
	negativeSumCount := 0
	total := 0.0
	for _, transaction := range transactions {
		total += transaction.Amount
		if transaction.Amount >= 0 {
			positiveSum += transaction.Amount
			positiveSumCount++
		} else {
			negativeSum += transaction.Amount
			negativeSumCount++
		}
	}

	// TODO: count balance value for each balance
	data := BalancePageTemplateData{
		Balance:        balance,
		Balances:       []*ent.Balance{balance},
		Categories:     categories,
		Transactions:   transactions,
		CategoryTitles: categoryTitles,
		Total:          total,
		TotalGain:      positiveSum,
		AverageGain:    positiveSum / float64(positiveSumCount),
		TotalSpent:     negativeSum,
		AverageSpent:   negativeSum / float64(negativeSumCount),
	}

	h.Template.ExecuteTemplate(ctx.Writer, "balancePage", data)
}

type BalancePageTemplateData struct {
	Balance        *ent.Balance
	Balances       []*ent.Balance
	Transactions   []*ent.Transaction
	Categories     []*ent.Category
	CategoryTitles map[int]string
	Total          float64
	TotalGain      float64
	TotalSpent     float64
	AverageGain    float64
	AverageSpent   float64
}
