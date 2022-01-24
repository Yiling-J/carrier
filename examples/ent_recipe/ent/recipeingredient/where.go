// Code generated by entc, DO NOT EDIT.

package recipeingredient

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// Unit applies equality check predicate on the "unit" field. It's identical to UnitEQ.
func Unit(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...float32) predicate.RecipeIngredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...float32) predicate.RecipeIngredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v float32) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// UnitEQ applies the EQ predicate on the "unit" field.
func UnitEQ(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnit), v))
	})
}

// UnitNEQ applies the NEQ predicate on the "unit" field.
func UnitNEQ(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnit), v))
	})
}

// UnitIn applies the In predicate on the "unit" field.
func UnitIn(vs ...string) predicate.RecipeIngredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnit), v...))
	})
}

// UnitNotIn applies the NotIn predicate on the "unit" field.
func UnitNotIn(vs ...string) predicate.RecipeIngredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnit), v...))
	})
}

// UnitGT applies the GT predicate on the "unit" field.
func UnitGT(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnit), v))
	})
}

// UnitGTE applies the GTE predicate on the "unit" field.
func UnitGTE(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnit), v))
	})
}

// UnitLT applies the LT predicate on the "unit" field.
func UnitLT(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnit), v))
	})
}

// UnitLTE applies the LTE predicate on the "unit" field.
func UnitLTE(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnit), v))
	})
}

// UnitContains applies the Contains predicate on the "unit" field.
func UnitContains(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnit), v))
	})
}

// UnitHasPrefix applies the HasPrefix predicate on the "unit" field.
func UnitHasPrefix(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnit), v))
	})
}

// UnitHasSuffix applies the HasSuffix predicate on the "unit" field.
func UnitHasSuffix(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnit), v))
	})
}

// UnitEqualFold applies the EqualFold predicate on the "unit" field.
func UnitEqualFold(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnit), v))
	})
}

// UnitContainsFold applies the ContainsFold predicate on the "unit" field.
func UnitContainsFold(v string) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnit), v))
	})
}

// HasIngredient applies the HasEdge predicate on the "ingredient" edge.
func HasIngredient() predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IngredientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IngredientTable, IngredientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIngredientWith applies the HasEdge predicate on the "ingredient" edge with a given conditions (other predicates).
func HasIngredientWith(preds ...predicate.Ingredient) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IngredientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IngredientTable, IngredientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RecipeIngredient) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RecipeIngredient) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RecipeIngredient) predicate.RecipeIngredient {
	return predicate.RecipeIngredient(func(s *sql.Selector) {
		p(s.Not())
	})
}