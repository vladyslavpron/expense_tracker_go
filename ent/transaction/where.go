// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"tracker/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDescription, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// BalanceID applies equality check predicate on the "balance_id" field. It's identical to BalanceIDEQ.
func BalanceID(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldBalanceID, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCategoryID, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldDescription, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float32) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldAmount, v))
}

// BalanceIDEQ applies the EQ predicate on the "balance_id" field.
func BalanceIDEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldBalanceID, v))
}

// BalanceIDNEQ applies the NEQ predicate on the "balance_id" field.
func BalanceIDNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldBalanceID, v))
}

// BalanceIDIn applies the In predicate on the "balance_id" field.
func BalanceIDIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldBalanceID, vs...))
}

// BalanceIDNotIn applies the NotIn predicate on the "balance_id" field.
func BalanceIDNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldBalanceID, vs...))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCategoryID, vs...))
}

// HasBalance applies the HasEdge predicate on the "balance" edge.
func HasBalance() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BalanceTable, BalanceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBalanceWith applies the HasEdge predicate on the "balance" edge with a given conditions (other predicates).
func HasBalanceWith(preds ...predicate.Balance) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newBalanceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.NotPredicates(p))
}
