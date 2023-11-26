package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Transaction struct {
	ent.Schema
}

// Fields of the User.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Default(""),
		field.Float32("amount"),
		field.Int("balance_id"),
		field.Int("category_id"),
	}
}

// Edges of the User.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("balance", Balance.Type).Ref("transactions").Field("balance_id").Unique().Required(),
		edge.From("category", Category.Type).Ref("transactions").Field("category_id").Unique().Required(),
	}
}
