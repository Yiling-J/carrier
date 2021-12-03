package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Step holds the schema definition for the Step entity.
type Step struct {
	ent.Schema
}

// Fields of the Step.
func (Step) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
	}
}

// Edges of the Step.
func (Step) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).Ref("steps").Unique(),
	}
}
