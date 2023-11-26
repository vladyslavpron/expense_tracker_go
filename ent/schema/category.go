package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Category struct {
	ent.Schema
}

// Fields of the User.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default(""),
	}
}

// Edges of the User.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("balance", Balance.Type).Ref("categories").Unique(),
		edge.To("transactions", Transaction.Type),
	}
}
