// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/ingredient"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/recipeingredient"
)

// RecipeIngredient is the model entity for the RecipeIngredient schema.
type RecipeIngredient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity float32 `json:"quantity,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit string `json:"unit,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecipeIngredientQuery when eager-loading is set.
	Edges                         RecipeIngredientEdges `json:"edges"`
	ingredient_recipe_ingredients *int
	recipe_recipe_ingredients     *int
}

// RecipeIngredientEdges holds the relations/edges for other nodes in the graph.
type RecipeIngredientEdges struct {
	// Ingredient holds the value of the ingredient edge.
	Ingredient *Ingredient `json:"ingredient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// IngredientOrErr returns the Ingredient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecipeIngredientEdges) IngredientOrErr() (*Ingredient, error) {
	if e.loadedTypes[0] {
		if e.Ingredient == nil {
			// The edge ingredient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ingredient.Label}
		}
		return e.Ingredient, nil
	}
	return nil, &NotLoadedError{edge: "ingredient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RecipeIngredient) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case recipeingredient.FieldQuantity:
			values[i] = new(sql.NullFloat64)
		case recipeingredient.FieldID:
			values[i] = new(sql.NullInt64)
		case recipeingredient.FieldUnit:
			values[i] = new(sql.NullString)
		case recipeingredient.ForeignKeys[0]: // ingredient_recipe_ingredients
			values[i] = new(sql.NullInt64)
		case recipeingredient.ForeignKeys[1]: // recipe_recipe_ingredients
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RecipeIngredient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RecipeIngredient fields.
func (ri *RecipeIngredient) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recipeingredient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ri.ID = int(value.Int64)
		case recipeingredient.FieldQuantity:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				ri.Quantity = float32(value.Float64)
			}
		case recipeingredient.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				ri.Unit = value.String
			}
		case recipeingredient.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field ingredient_recipe_ingredients", value)
			} else if value.Valid {
				ri.ingredient_recipe_ingredients = new(int)
				*ri.ingredient_recipe_ingredients = int(value.Int64)
			}
		case recipeingredient.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field recipe_recipe_ingredients", value)
			} else if value.Valid {
				ri.recipe_recipe_ingredients = new(int)
				*ri.recipe_recipe_ingredients = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryIngredient queries the "ingredient" edge of the RecipeIngredient entity.
func (ri *RecipeIngredient) QueryIngredient() *IngredientQuery {
	return (&RecipeIngredientClient{config: ri.config}).QueryIngredient(ri)
}

// Update returns a builder for updating this RecipeIngredient.
// Note that you need to call RecipeIngredient.Unwrap() before calling this method if this RecipeIngredient
// was returned from a transaction, and the transaction was committed or rolled back.
func (ri *RecipeIngredient) Update() *RecipeIngredientUpdateOne {
	return (&RecipeIngredientClient{config: ri.config}).UpdateOne(ri)
}

// Unwrap unwraps the RecipeIngredient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ri *RecipeIngredient) Unwrap() *RecipeIngredient {
	tx, ok := ri.config.driver.(*txDriver)
	if !ok {
		panic("ent: RecipeIngredient is not a transactional entity")
	}
	ri.config.driver = tx.drv
	return ri
}

// String implements the fmt.Stringer.
func (ri *RecipeIngredient) String() string {
	var builder strings.Builder
	builder.WriteString("RecipeIngredient(")
	builder.WriteString(fmt.Sprintf("id=%v", ri.ID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", ri.Quantity))
	builder.WriteString(", unit=")
	builder.WriteString(ri.Unit)
	builder.WriteByte(')')
	return builder.String()
}

// RecipeIngredients is a parsable slice of RecipeIngredient.
type RecipeIngredients []*RecipeIngredient

func (ri RecipeIngredients) config(cfg config) {
	for _i := range ri {
		ri[_i].config = cfg
	}
}