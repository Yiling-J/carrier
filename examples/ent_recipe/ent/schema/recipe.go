package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Recipe holds the schema definition for the Recipe entity.
type Recipe struct {
	ent.Schema
}

// Fields of the Recipe.
func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("servings"),
	}
}

// Edges of the Recipe.
func (Recipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("recipes").Unique(),
		edge.To("tags", Category.Type),
		edge.To("steps", Step.Type),
		edge.To("recipe_ingredients", RecipeIngredient.Type),
	}
}
