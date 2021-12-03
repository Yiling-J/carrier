package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent"
)

var (
	Schemas = []carrier.Schema{
		&carrier.EntSchema{
			To: &ent.RecipeCreate{},
			Posts: []carrier.PostField{
				{Name: "stepsCount", Input: 1},
				{Name: "ingredientsCount", Input: 1},
			},
			Traits: []string{"vegan", "keto"},
		},
		&carrier.EntSchema{
			To:    &ent.UserCreate{},
			Posts: []carrier.PostField{{Name: "recipes", Input: 1}},
		},
		&carrier.EntSchema{To: &ent.IngredientCreate{}},
		&carrier.EntSchema{To: &ent.CategoryCreate{}},
		&carrier.EntSchema{To: &ent.RecipeIngredientCreate{}},
		&carrier.EntSchema{To: &ent.StepCreate{}},
	}
)
