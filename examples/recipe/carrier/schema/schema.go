package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/Yiling-J/carrier/examples/recipe/model"
)

var (
	Schemas = []carrier.Schema{
		&carrier.StructSchema{
			To: model.Recipe{},
			Posts: []carrier.PostField{
				{Name: "stepsCount", Input: 1},
				{Name: "ingredientsCount", Input: 1},
			},
			Traits: []string{"vegan", "keto"},
		},
		&carrier.StructSchema{To: model.RecipeStep{}},
		&carrier.StructSchema{To: model.Ingredient{}},
		&carrier.StructSchema{To: model.RecipeIngredient{}},
		&carrier.StructSchema{To: model.Category{}},
		&carrier.StructSchema{
			To:    model.User{},
			Posts: []carrier.PostField{{Name: "recipes", Input: 1}},
		},
	}
)
