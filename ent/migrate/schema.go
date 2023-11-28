// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BalancesColumns holds the columns for the "balances" table.
	BalancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Default: ""},
	}
	// BalancesTable holds the schema information for the "balances" table.
	BalancesTable = &schema.Table{
		Name:       "balances",
		Columns:    BalancesColumns,
		PrimaryKey: []*schema.Column{BalancesColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "balance_id", Type: field.TypeInt},
		{Name: "category_id", Type: field.TypeInt},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_balances_transactions",
				Columns:    []*schema.Column{TransactionsColumns[4]},
				RefColumns: []*schema.Column{BalancesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "transactions_categories_transactions",
				Columns:    []*schema.Column{TransactionsColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BalancesTable,
		CategoriesTable,
		TransactionsTable,
		UsersTable,
	}
)

func init() {
	TransactionsTable.ForeignKeys[0].RefTable = BalancesTable
	TransactionsTable.ForeignKeys[1].RefTable = CategoriesTable
}
