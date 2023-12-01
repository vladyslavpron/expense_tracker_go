package viewHandler

import (
	"tracker/core/balance"
	"tracker/core/balance/currency"
	transactionHandler "tracker/core/handlers/transaction"
	"tracker/ent"

	"github.com/gin-gonic/gin"
)

func (h *ViewHandler) Main(ctx *gin.Context) {

	balances, _ := h.BalanceHandler.GetBalances(ctx)
	categories, _ := h.CategoryHandler.GetCategories(ctx)
	transactions, _ := h.TransactionHandler.GetTransactions(ctx, transactionHandler.GetTransactionsParams{})

	balancesTransactionsSumCount, _ := balance.CountBalancesTransactions(ctx, h.BalanceHandler.DB)

	categoryTitles := make(map[int]string, len(categories))
	for _, category := range categories {
		categoryTitles[category.ID] = category.Title
	}

	// TODO: count balance value for each balance
	data := MainTemplateData{
		Balances:                     balancesToModel(balances, balancesTransactionsSumCount),
		Categories:                   categories,
		Transactions:                 transactions,
		CategoryTitles:               categoryTitles,
		BalancesTransactionsSumCount: balancesTransactionsSumCount,
	}

	h.Template.ExecuteTemplate(ctx.Writer, "main", data)
}

type MainTemplateData struct {
	Balances                     []Balance
	Categories                   []*ent.Category
	Transactions                 []*ent.Transaction
	CategoryTitles               map[int]string
	BalancesTransactionsSumCount []balance.TransactionSumCount
}

func balancesToModel(balances []*ent.Balance, stats []balance.TransactionSumCount) (r []Balance) {
	// map stats to map
	statsMap := make(map[int]balance.TransactionSumCount)

	for _, s := range stats {
		statsMap[s.BalanceID] = s
	}

	for _, b := range balances {
		stat := statsMap[b.ID]
		r = append(r, Balance{ID: b.ID, Title: b.Title, Currency: b.Currency, UsdToCurrency: b.UsdToCurrency, TransactionsCount: stat.Count, Total: stat.Sum})
	}
	return
}

type Balance struct {
	ID                int
	Title             string
	Currency          currency.Currency
	UsdToCurrency     float64
	TransactionsCount int
	Total             float64
}
