package viewHandler

import (
	transactionHandler "tracker/core/handlers/transaction"
	"tracker/ent"

	"github.com/gin-gonic/gin"
)

func (h *ViewHandler) Main(ctx *gin.Context) {

	balances, _ := h.BalanceHandler.GetBalances(ctx)
	categories, _ := h.CategoryHandler.GetCategories(ctx)
	transactions, _ := h.TransactionHandler.GetTransactions(ctx, transactionHandler.GetTransactionsParams{})

	categoryTitles := make(map[int]string, len(categories))
	for _, category := range categories {
		categoryTitles[category.ID] = category.Title
	}

	// TODO: count balance value for each balance
	data := MainTemplateData{
		Balances:       balances,
		Categories:     categories,
		Transactions:   transactions,
		CategoryTitles: categoryTitles,
	}

	h.Template.ExecuteTemplate(ctx.Writer, "main", data)
}

type MainTemplateData struct {
	Balances       []*ent.Balance
	Categories     []*ent.Category
	Transactions   []*ent.Transaction
	CategoryTitles map[int]string
}
