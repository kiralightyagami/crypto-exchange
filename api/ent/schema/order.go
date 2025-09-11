package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("market").
			Default("unknown"),
		field.Int("price").
			Default(0),
		field.Int("quantity").
			Default(0),
		field.Enum("order_type").
			Values("BUY", "SELL"),
		field.Enum("order_status").
			Values("PENDING", "FILLED", "CANCELLED"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("filled_at").
			Optional(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("orders").Unique(),
	}
}
