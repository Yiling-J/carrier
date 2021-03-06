// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/examples/recipe/model"

	"context"
)

type recipeStepMutation struct {
	textType int
	textFunc func(ctx context.Context, i *model.RecipeStep, c int) error

	beforeCreateFunc func(ctx context.Context, i *model.RecipeStep) error
	afterCreateFunc  func(ctx context.Context, i *model.RecipeStep) error
}
type RecipeStepMetaFactory struct {
	mutation recipeStepMutation
}
type recipeStepTrait struct {
	mutation recipeStepMutation
	updates  []func(m *recipeStepMutation)
}

func RecipeStepTrait() *recipeStepTrait {
	return &recipeStepTrait{}
}
func (*recipeStepMutation) beforeCreateMutateFunc(fn func(ctx context.Context, i *model.RecipeStep) error) func(m *recipeStepMutation) {
	return func(m *recipeStepMutation) {
		m.beforeCreateFunc = fn
	}
}
func (*recipeStepMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *model.RecipeStep) error) func(m *recipeStepMutation) {
	return func(m *recipeStepMutation) {
		m.afterCreateFunc = fn
	}
}

func (*recipeStepMutation) textSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *recipeStepMutation) {
	return func(m *recipeStepMutation) {
		m.textType = TypeSequence
		m.textFunc = func(ctx context.Context, i *model.RecipeStep, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Text = value
			return nil
		}
	}
}
func (*recipeStepMutation) textLazyMutateFunc(fn func(ctx context.Context, i *model.RecipeStep) (string, error)) func(m *recipeStepMutation) {
	return func(m *recipeStepMutation) {
		m.textType = TypeLazy
		m.textFunc = func(ctx context.Context, i *model.RecipeStep, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Text = value
			return nil
		}
	}
}
func (*recipeStepMutation) textDefaultMutateFunc(v string) func(m *recipeStepMutation) {
	return func(m *recipeStepMutation) {
		m.textType = TypeDefault
		m.textFunc = func(ctx context.Context, i *model.RecipeStep, c int) error {

			i.Text = v
			return nil
		}
	}
}
func (*recipeStepMutation) textFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *recipeStepMutation) {
	return func(m *recipeStepMutation) {
		m.textType = TypeFactory
		m.textFunc = func(ctx context.Context, i *model.RecipeStep, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Text = value

			return nil
		}
	}
}

// SetTextSequence register a function which accept a sequence counter and set return value to Text field
func (f *RecipeStepMetaFactory) SetTextSequence(fn func(ctx context.Context, i int) (string, error)) *RecipeStepMetaFactory {
	f.mutation.textSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetTextLazy register a function which accept the build struct and set return value to Text field
func (f *RecipeStepMetaFactory) SetTextLazy(fn func(ctx context.Context, i *model.RecipeStep) (string, error)) *RecipeStepMetaFactory {
	f.mutation.textLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetTextDefault assign a default value to Text field
func (f *RecipeStepMetaFactory) SetTextDefault(v string) *RecipeStepMetaFactory {
	f.mutation.textDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetTextFactory register a factory function and assign return value to Text, you can also use related factory's Create/CreateV as input function here
func (f *RecipeStepMetaFactory) SetTextFactory(fn func(ctx context.Context) (string, error)) *RecipeStepMetaFactory {
	f.mutation.textFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetTextSequence register a function which accept a sequence counter and set return value to Text field
func (t *recipeStepTrait) SetTextSequence(fn func(ctx context.Context, i int) (string, error)) *recipeStepTrait {
	t.updates = append(t.updates, t.mutation.textSequenceMutateFunc(fn))
	return t
}

// SetTextLazy register a function which accept the build struct and set return value to Text field
func (t *recipeStepTrait) SetTextLazy(fn func(ctx context.Context, i *model.RecipeStep) (string, error)) *recipeStepTrait {
	t.updates = append(t.updates, t.mutation.textLazyMutateFunc(fn))
	return t
}

// SetTextDefault assign a default value to Text field
func (t *recipeStepTrait) SetTextDefault(v string) *recipeStepTrait {
	t.updates = append(t.updates, t.mutation.textDefaultMutateFunc(v))
	return t
}

// SetTextFactory register a factory function and assign return value to Text, you can also use related factory's Create/CreateV as input function here
func (t *recipeStepTrait) SetTextFactory(fn func(ctx context.Context) (string, error)) *recipeStepTrait {
	t.updates = append(t.updates, t.mutation.textFactoryMutateFunc(fn))
	return t
}

// SetAfterCreateFunc register a function to be called after struct create
func (f *RecipeStepMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *model.RecipeStep) error) *RecipeStepMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

// SetBeforeCreateFunc register a function to be called before struct create
func (f *RecipeStepMetaFactory) SetBeforeCreateFunc(fn func(ctx context.Context, i *model.RecipeStep) error) *RecipeStepMetaFactory {
	f.mutation.beforeCreateFunc = fn
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (t *recipeStepTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *model.RecipeStep) error) *recipeStepTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

// SetBeforeCreateFunc register a function to be called before struct create
func (t *recipeStepTrait) SetBeforeCreateFunc(fn func(ctx context.Context, i *model.RecipeStep) error) *recipeStepTrait {
	t.updates = append(t.updates, t.mutation.beforeCreateMutateFunc(fn))
	return t
}

// Build create a  RecipeStepFactory from RecipeStepMetaFactory
func (f *RecipeStepMetaFactory) Build() *RecipeStepFactory {
	return &RecipeStepFactory{meta: *f, counter: &Counter{}}
}

type RecipeStepFactory struct {
	meta    RecipeStepMetaFactory
	counter *Counter
}

// SetText set the Text field
func (f *RecipeStepFactory) SetText(i string) *RecipeStepBuilder {
	builder := &RecipeStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetText(i)

	return builder
}

// Create return a new *model.RecipeStep
func (f *RecipeStepFactory) Create(ctx context.Context) (*model.RecipeStep, error) {
	builder := &RecipeStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.Create(ctx)
}

// CreateV return a new model.RecipeStep
func (f *RecipeStepFactory) CreateV(ctx context.Context) (model.RecipeStep, error) {
	builder := &RecipeStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateV(ctx)
}

// CreateBatch return a []*model.RecipeStep slice
func (f *RecipeStepFactory) CreateBatch(ctx context.Context, n int) ([]*model.RecipeStep, error) {
	builder := &RecipeStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatch(ctx, n)
}

// CreateBatchV return a []model.RecipeStep slice
func (f *RecipeStepFactory) CreateBatchV(ctx context.Context, n int) ([]model.RecipeStep, error) {
	builder := &RecipeStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatchV(ctx, n)
}

type RecipeStepBuilder struct {
	factory  *RecipeStepFactory
	mutation recipeStepMutation
	counter  *Counter

	textOverride  string
	textOverriden bool
}

// SetText set the Text field
func (b *RecipeStepBuilder) SetText(i string) *RecipeStepBuilder {
	b.textOverride = i
	b.textOverriden = true
	return b
}

// CreateV return a new model.RecipeStep
func (b *RecipeStepBuilder) CreateV(ctx context.Context) (model.RecipeStep, error) {
	var d model.RecipeStep
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

// Create return a new *model.RecipeStep
func (b *RecipeStepBuilder) Create(ctx context.Context) (*model.RecipeStep, error) {

	var preSlice = []func(ctx context.Context, i *model.RecipeStep, c int) error{}
	var lazySlice = []func(ctx context.Context, i *model.RecipeStep, c int) error{}
	var postSlice = []func(ctx context.Context, i *model.RecipeStep, c int) error{}

	index := b.counter.Get()
	_ = index

	if b.textOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.RecipeStep, c int) error {
			value := b.textOverride

			i.Text = value
			return nil
		})
	} else {
		switch b.mutation.textType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.textFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.textFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.textFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.textFunc)
		}
	}

	v := &model.RecipeStep{}

	for _, f := range preSlice {

		err := f(ctx, v, index)

		if err != nil {
			return nil, err
		}
	}
	for _, f := range lazySlice {

		err := f(ctx, v, index)

		if err != nil {
			return nil, err
		}
	}
	if b.mutation.beforeCreateFunc != nil {
		if err := b.mutation.beforeCreateFunc(ctx, v); err != nil {
			return nil, err
		}
	}

	new := v

	if b.mutation.afterCreateFunc != nil {
		err := b.mutation.afterCreateFunc(ctx, new)
		if err != nil {
			return nil, err
		}
	}
	for _, f := range postSlice {
		err := f(ctx, new, index)
		if err != nil {
			return nil, err
		}
	}

	return new, nil
}
func (b *RecipeStepBuilder) CreateBatch(ctx context.Context, n int) ([]*model.RecipeStep, error) {
	var results []*model.RecipeStep
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *RecipeStepBuilder) CreateBatchV(ctx context.Context, n int) ([]model.RecipeStep, error) {
	var results []model.RecipeStep
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
