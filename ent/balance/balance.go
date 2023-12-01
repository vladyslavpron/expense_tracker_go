// Code generated by ent, DO NOT EDIT.

package balance

import (
	"fmt"
	"tracker/core/balance/currency"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the balance type in the database.
	Label = "balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldUsdToCurrency holds the string denoting the usd_to_currency field in the database.
	FieldUsdToCurrency = "usd_to_currency"
	// EdgeTransactions holds the string denoting the transactions edge name in mutations.
	EdgeTransactions = "transactions"
	// Table holds the table name of the balance in the database.
	Table = "balances"
	// TransactionsTable is the table that holds the transactions relation/edge.
	TransactionsTable = "transactions"
	// TransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionsInverseTable = "transactions"
	// TransactionsColumn is the table column denoting the transactions relation/edge.
	TransactionsColumn = "balance_id"
)

// Columns holds all SQL columns for balance fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldCurrency,
	FieldUsdToCurrency,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultUsdToCurrency holds the default value on creation for the "usd_to_currency" field.
	DefaultUsdToCurrency float64
)

const DefaultCurrency currency.Currency = "UAH"

// CurrencyValidator is a validator for the "currency" field enum values. It is called by the builders before save.
func CurrencyValidator(c currency.Currency) error {
	switch c {
	case "USD", "EUR", "UAH":
		return nil
	default:
		return fmt.Errorf("balance: invalid enum value for currency field: %q", c)
	}
}

// OrderOption defines the ordering options for the Balance queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByCurrency orders the results by the currency field.
func ByCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrency, opts...).ToFunc()
}

// ByUsdToCurrency orders the results by the usd_to_currency field.
func ByUsdToCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsdToCurrency, opts...).ToFunc()
}

// ByTransactionsCount orders the results by transactions count.
func ByTransactionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTransactionsStep(), opts...)
	}
}

// ByTransactions orders the results by transactions terms.
func ByTransactions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTransactionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTransactionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TransactionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
	)
}
