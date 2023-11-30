package schema

import (
	"tracker/core/balance"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Balance struct {
	ent.Schema
}

// Fields of the User.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default(""),
		field.Enum("currency").GoType(balance.Currency("")).Default(string(balance.UAH)),
		field.Float("usd_to_currency").Default(1.0),
	}
}

// Edges of the User.
func (Balance) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("categories", Category.Type),
		edge.To("transactions", Transaction.Type),
	}
}

func (Balance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
