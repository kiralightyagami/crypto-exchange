package schema

import (
	"entgo.io/ent"
)

// Klines holds the schema definition for the Klines entity.
type Klines struct {
	ent.Schema
}

// Fields of the Klines.
func (Klines) Fields() []ent.Field {
	return nil
}

// Edges of the Klines.
func (Klines) Edges() []ent.Edge {
	return nil
}
