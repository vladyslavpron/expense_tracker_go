package viewHandler

import (
	"html/template"
	balanceHandler "tracker/core/handlers/balance"
	categoryHandler "tracker/core/handlers/category"
	transactionHandler "tracker/core/handlers/transaction"
)

type ViewHandler struct {
	BalanceHandler     *balanceHandler.BalanceHandler
	CategoryHandler    *categoryHandler.CategoryHandler
	TransactionHandler *transactionHandler.TransactionHandler
	Template           *template.Template
}
