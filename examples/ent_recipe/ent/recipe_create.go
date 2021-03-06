// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/category"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/recipe"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/recipeingredient"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/step"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/user"
)

// RecipeCreate is the builder for creating a Recipe entity.
type RecipeCreate struct {
	config
	mutation *RecipeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *RecipeCreate) SetName(s string) *RecipeCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetServings sets the "servings" field.
func (rc *RecipeCreate) SetServings(i int) *RecipeCreate {
	rc.mutation.SetServings(i)
	return rc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (rc *RecipeCreate) SetAuthorID(id int) *RecipeCreate {
	rc.mutation.SetAuthorID(id)
	return rc
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (rc *RecipeCreate) SetNillableAuthorID(id *int) *RecipeCreate {
	if id != nil {
		rc = rc.SetAuthorID(*id)
	}
	return rc
}

// SetAuthor sets the "author" edge to the User entity.
func (rc *RecipeCreate) SetAuthor(u *User) *RecipeCreate {
	return rc.SetAuthorID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Category entity by IDs.
func (rc *RecipeCreate) AddTagIDs(ids ...int) *RecipeCreate {
	rc.mutation.AddTagIDs(ids...)
	return rc
}

// AddTags adds the "tags" edges to the Category entity.
func (rc *RecipeCreate) AddTags(c ...*Category) *RecipeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddTagIDs(ids...)
}

// AddStepIDs adds the "steps" edge to the Step entity by IDs.
func (rc *RecipeCreate) AddStepIDs(ids ...int) *RecipeCreate {
	rc.mutation.AddStepIDs(ids...)
	return rc
}

// AddSteps adds the "steps" edges to the Step entity.
func (rc *RecipeCreate) AddSteps(s ...*Step) *RecipeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddStepIDs(ids...)
}

// AddRecipeIngredientIDs adds the "recipe_ingredients" edge to the RecipeIngredient entity by IDs.
func (rc *RecipeCreate) AddRecipeIngredientIDs(ids ...int) *RecipeCreate {
	rc.mutation.AddRecipeIngredientIDs(ids...)
	return rc
}

// AddRecipeIngredients adds the "recipe_ingredients" edges to the RecipeIngredient entity.
func (rc *RecipeCreate) AddRecipeIngredients(r ...*RecipeIngredient) *RecipeCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRecipeIngredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (rc *RecipeCreate) Mutation() *RecipeMutation {
	return rc.mutation
}

// Save creates the Recipe in the database.
func (rc *RecipeCreate) Save(ctx context.Context) (*Recipe, error) {
	var (
		err  error
		node *Recipe
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecipeCreate) SaveX(ctx context.Context) *Recipe {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecipeCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecipeCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecipeCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := rc.mutation.Servings(); !ok {
		return &ValidationError{Name: "servings", err: errors.New(`ent: missing required field "servings"`)}
	}
	return nil
}

func (rc *RecipeCreate) sqlSave(ctx context.Context) (*Recipe, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RecipeCreate) createSpec() (*Recipe, *sqlgraph.CreateSpec) {
	var (
		_node = &Recipe{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recipe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipe.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rc.mutation.Servings(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldServings,
		})
		_node.Servings = value
	}
	if nodes := rc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recipe.AuthorTable,
			Columns: []string{recipe.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_recipes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RecipeIngredientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RecipeCreateBulk is the builder for creating many Recipe entities in bulk.
type RecipeCreateBulk struct {
	config
	builders []*RecipeCreate
}

// Save creates the Recipe entities in the database.
func (rcb *RecipeCreateBulk) Save(ctx context.Context) ([]*Recipe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Recipe, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecipeCreateBulk) SaveX(ctx context.Context) []*Recipe {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecipeCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecipeCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
