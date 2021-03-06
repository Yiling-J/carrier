// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent"

	"context"
)

type EntStepMutator struct {
	Recipe *ent.Recipe

	RecipeID int

	Text string

	_creator *ent.StepCreate
}

func (m *EntStepMutator) EntCreator() *ent.StepCreate {
	return m._creator
}

type entStepMutation struct {
	recipeType int
	recipeFunc func(ctx context.Context, i *EntStepMutator, c int) error

	recipeIDType int
	recipeIDFunc func(ctx context.Context, i *EntStepMutator, c int) error

	textType int
	textFunc func(ctx context.Context, i *EntStepMutator, c int) error

	beforeCreateFunc func(ctx context.Context, i *EntStepMutator) error
	afterCreateFunc  func(ctx context.Context, i *ent.Step) error
}
type EntStepMetaFactory struct {
	mutation entStepMutation
}
type entStepTrait struct {
	mutation entStepMutation
	updates  []func(m *entStepMutation)
}

func EntStepTrait() *entStepTrait {
	return &entStepTrait{}
}
func (*entStepMutation) beforeCreateMutateFunc(fn func(ctx context.Context, i *EntStepMutator) error) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.beforeCreateFunc = fn
	}
}
func (*entStepMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.Step) error) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entStepMutation) recipeSequenceMutateFunc(fn func(ctx context.Context, i int) (*ent.Recipe, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeType = TypeSequence
		m.recipeFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetRecipe(value)

			i.Recipe = value
			return nil
		}
	}
}
func (*entStepMutation) recipeLazyMutateFunc(fn func(ctx context.Context, i *EntStepMutator) (*ent.Recipe, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeType = TypeLazy
		m.recipeFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetRecipe(value)

			i.Recipe = value
			return nil
		}
	}
}
func (*entStepMutation) recipeDefaultMutateFunc(v *ent.Recipe) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeType = TypeDefault
		m.recipeFunc = func(ctx context.Context, i *EntStepMutator, c int) error {

			i.EntCreator().SetRecipe(v)

			i.Recipe = v
			return nil
		}
	}
}
func (*entStepMutation) recipeFactoryMutateFunc(fn func(ctx context.Context) (*ent.Recipe, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeType = TypeFactory
		m.recipeFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetRecipe(value)

			i.Recipe = value

			return nil
		}
	}
}

// SetRecipeSequence register a function which accept a sequence counter and set return value to Recipe field
func (f *EntStepMetaFactory) SetRecipeSequence(fn func(ctx context.Context, i int) (*ent.Recipe, error)) *EntStepMetaFactory {
	f.mutation.recipeSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetRecipeLazy register a function which accept the build struct and set return value to Recipe field
func (f *EntStepMetaFactory) SetRecipeLazy(fn func(ctx context.Context, i *EntStepMutator) (*ent.Recipe, error)) *EntStepMetaFactory {
	f.mutation.recipeLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetRecipeDefault assign a default value to Recipe field
func (f *EntStepMetaFactory) SetRecipeDefault(v *ent.Recipe) *EntStepMetaFactory {
	f.mutation.recipeDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetRecipeFactory register a factory function and assign return value to Recipe, you can also use related factory's Create/CreateV as input function here
func (f *EntStepMetaFactory) SetRecipeFactory(fn func(ctx context.Context) (*ent.Recipe, error)) *EntStepMetaFactory {
	f.mutation.recipeFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetRecipeSequence register a function which accept a sequence counter and set return value to Recipe field
func (t *entStepTrait) SetRecipeSequence(fn func(ctx context.Context, i int) (*ent.Recipe, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeSequenceMutateFunc(fn))
	return t
}

// SetRecipeLazy register a function which accept the build struct and set return value to Recipe field
func (t *entStepTrait) SetRecipeLazy(fn func(ctx context.Context, i *EntStepMutator) (*ent.Recipe, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeLazyMutateFunc(fn))
	return t
}

// SetRecipeDefault assign a default value to Recipe field
func (t *entStepTrait) SetRecipeDefault(v *ent.Recipe) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeDefaultMutateFunc(v))
	return t
}

// SetRecipeFactory register a factory function and assign return value to Recipe, you can also use related factory's Create/CreateV as input function here
func (t *entStepTrait) SetRecipeFactory(fn func(ctx context.Context) (*ent.Recipe, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeFactoryMutateFunc(fn))
	return t
}

func (*entStepMutation) recipeIDSequenceMutateFunc(fn func(ctx context.Context, i int) (int, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeIDType = TypeSequence
		m.recipeIDFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetRecipeID(value)

			i.RecipeID = value
			return nil
		}
	}
}
func (*entStepMutation) recipeIDLazyMutateFunc(fn func(ctx context.Context, i *EntStepMutator) (int, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeIDType = TypeLazy
		m.recipeIDFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetRecipeID(value)

			i.RecipeID = value
			return nil
		}
	}
}
func (*entStepMutation) recipeIDDefaultMutateFunc(v int) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeIDType = TypeDefault
		m.recipeIDFunc = func(ctx context.Context, i *EntStepMutator, c int) error {

			i.EntCreator().SetRecipeID(v)

			i.RecipeID = v
			return nil
		}
	}
}
func (*entStepMutation) recipeIDFactoryMutateFunc(fn func(ctx context.Context) (int, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.recipeIDType = TypeFactory
		m.recipeIDFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetRecipeID(value)

			i.RecipeID = value

			return nil
		}
	}
}

// SetRecipeIDSequence register a function which accept a sequence counter and set return value to RecipeID field
func (f *EntStepMetaFactory) SetRecipeIDSequence(fn func(ctx context.Context, i int) (int, error)) *EntStepMetaFactory {
	f.mutation.recipeIDSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetRecipeIDLazy register a function which accept the build struct and set return value to RecipeID field
func (f *EntStepMetaFactory) SetRecipeIDLazy(fn func(ctx context.Context, i *EntStepMutator) (int, error)) *EntStepMetaFactory {
	f.mutation.recipeIDLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetRecipeIDDefault assign a default value to RecipeID field
func (f *EntStepMetaFactory) SetRecipeIDDefault(v int) *EntStepMetaFactory {
	f.mutation.recipeIDDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetRecipeIDFactory register a factory function and assign return value to RecipeID, you can also use related factory's Create/CreateV as input function here
func (f *EntStepMetaFactory) SetRecipeIDFactory(fn func(ctx context.Context) (int, error)) *EntStepMetaFactory {
	f.mutation.recipeIDFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetRecipeIDSequence register a function which accept a sequence counter and set return value to RecipeID field
func (t *entStepTrait) SetRecipeIDSequence(fn func(ctx context.Context, i int) (int, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeIDSequenceMutateFunc(fn))
	return t
}

// SetRecipeIDLazy register a function which accept the build struct and set return value to RecipeID field
func (t *entStepTrait) SetRecipeIDLazy(fn func(ctx context.Context, i *EntStepMutator) (int, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeIDLazyMutateFunc(fn))
	return t
}

// SetRecipeIDDefault assign a default value to RecipeID field
func (t *entStepTrait) SetRecipeIDDefault(v int) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeIDDefaultMutateFunc(v))
	return t
}

// SetRecipeIDFactory register a factory function and assign return value to RecipeID, you can also use related factory's Create/CreateV as input function here
func (t *entStepTrait) SetRecipeIDFactory(fn func(ctx context.Context) (int, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.recipeIDFactoryMutateFunc(fn))
	return t
}

func (*entStepMutation) textSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.textType = TypeSequence
		m.textFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetText(value)

			i.Text = value
			return nil
		}
	}
}
func (*entStepMutation) textLazyMutateFunc(fn func(ctx context.Context, i *EntStepMutator) (string, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.textType = TypeLazy
		m.textFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetText(value)

			i.Text = value
			return nil
		}
	}
}
func (*entStepMutation) textDefaultMutateFunc(v string) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.textType = TypeDefault
		m.textFunc = func(ctx context.Context, i *EntStepMutator, c int) error {

			i.EntCreator().SetText(v)

			i.Text = v
			return nil
		}
	}
}
func (*entStepMutation) textFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entStepMutation) {
	return func(m *entStepMutation) {
		m.textType = TypeFactory
		m.textFunc = func(ctx context.Context, i *EntStepMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetText(value)

			i.Text = value

			return nil
		}
	}
}

// SetTextSequence register a function which accept a sequence counter and set return value to Text field
func (f *EntStepMetaFactory) SetTextSequence(fn func(ctx context.Context, i int) (string, error)) *EntStepMetaFactory {
	f.mutation.textSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetTextLazy register a function which accept the build struct and set return value to Text field
func (f *EntStepMetaFactory) SetTextLazy(fn func(ctx context.Context, i *EntStepMutator) (string, error)) *EntStepMetaFactory {
	f.mutation.textLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetTextDefault assign a default value to Text field
func (f *EntStepMetaFactory) SetTextDefault(v string) *EntStepMetaFactory {
	f.mutation.textDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetTextFactory register a factory function and assign return value to Text, you can also use related factory's Create/CreateV as input function here
func (f *EntStepMetaFactory) SetTextFactory(fn func(ctx context.Context) (string, error)) *EntStepMetaFactory {
	f.mutation.textFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetTextSequence register a function which accept a sequence counter and set return value to Text field
func (t *entStepTrait) SetTextSequence(fn func(ctx context.Context, i int) (string, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.textSequenceMutateFunc(fn))
	return t
}

// SetTextLazy register a function which accept the build struct and set return value to Text field
func (t *entStepTrait) SetTextLazy(fn func(ctx context.Context, i *EntStepMutator) (string, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.textLazyMutateFunc(fn))
	return t
}

// SetTextDefault assign a default value to Text field
func (t *entStepTrait) SetTextDefault(v string) *entStepTrait {
	t.updates = append(t.updates, t.mutation.textDefaultMutateFunc(v))
	return t
}

// SetTextFactory register a factory function and assign return value to Text, you can also use related factory's Create/CreateV as input function here
func (t *entStepTrait) SetTextFactory(fn func(ctx context.Context) (string, error)) *entStepTrait {
	t.updates = append(t.updates, t.mutation.textFactoryMutateFunc(fn))
	return t
}

// SetAfterCreateFunc register a function to be called after struct create
func (f *EntStepMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Step) error) *EntStepMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

// SetBeforeCreateFunc register a function to be called before struct create
func (f *EntStepMetaFactory) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntStepMutator) error) *EntStepMetaFactory {
	f.mutation.beforeCreateFunc = fn
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (t *entStepTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Step) error) *entStepTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

// SetBeforeCreateFunc register a function to be called before struct create
func (t *entStepTrait) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntStepMutator) error) *entStepTrait {
	t.updates = append(t.updates, t.mutation.beforeCreateMutateFunc(fn))
	return t
}

// Build create a  EntStepFactory from EntStepMetaFactory
func (f *EntStepMetaFactory) Build() *EntStepFactory {
	return &EntStepFactory{meta: *f, counter: &Counter{}}
}

type EntStepFactory struct {
	meta    EntStepMetaFactory
	counter *Counter

	client *ent.Client
}

// SetRecipe set the Recipe field
func (f *EntStepFactory) SetRecipe(i *ent.Recipe) *EntStepBuilder {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetRecipe(i)

	builder.client = f.client

	return builder
}

// SetRecipeID set the RecipeID field
func (f *EntStepFactory) SetRecipeID(i int) *EntStepBuilder {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetRecipeID(i)

	builder.client = f.client

	return builder
}

// SetText set the Text field
func (f *EntStepFactory) SetText(i string) *EntStepBuilder {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetText(i)

	builder.client = f.client

	return builder
}

// Create return a new *ent.Step
func (f *EntStepFactory) Create(ctx context.Context) (*ent.Step, error) {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.Create(ctx)
}

// CreateV return a new ent.Step
func (f *EntStepFactory) CreateV(ctx context.Context) (ent.Step, error) {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateV(ctx)
}

// CreateBatch return a []*ent.Step slice
func (f *EntStepFactory) CreateBatch(ctx context.Context, n int) ([]*ent.Step, error) {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatch(ctx, n)
}

// CreateBatchV return a []ent.Step slice
func (f *EntStepFactory) CreateBatchV(ctx context.Context, n int) ([]ent.Step, error) {
	builder := &EntStepBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatchV(ctx, n)
}

// Client set ent client to EntStepFactory
func (f *EntStepFactory) Client(c *ent.Client) *EntStepFactory {
	f.client = c
	return f
}

type EntStepBuilder struct {
	factory  *EntStepFactory
	mutation entStepMutation
	counter  *Counter

	recipeOverride  *ent.Recipe
	recipeOverriden bool

	recipeIDOverride  int
	recipeIDOverriden bool

	textOverride  string
	textOverriden bool

	client *ent.Client
}

func (b *EntStepBuilder) Client(c *ent.Client) *EntStepBuilder {
	b.client = c
	return b
}

// SetRecipe set the Recipe field
func (b *EntStepBuilder) SetRecipe(i *ent.Recipe) *EntStepBuilder {
	b.recipeOverride = i
	b.recipeOverriden = true
	return b
}

// SetRecipeID set the RecipeID field
func (b *EntStepBuilder) SetRecipeID(i int) *EntStepBuilder {
	b.recipeIDOverride = i
	b.recipeIDOverriden = true
	return b
}

// SetText set the Text field
func (b *EntStepBuilder) SetText(i string) *EntStepBuilder {
	b.textOverride = i
	b.textOverriden = true
	return b
}

// CreateV return a new ent.Step
func (b *EntStepBuilder) CreateV(ctx context.Context) (ent.Step, error) {
	var d ent.Step
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

// Create return a new *ent.Step
func (b *EntStepBuilder) Create(ctx context.Context) (*ent.Step, error) {

	var preSlice = []func(ctx context.Context, i *EntStepMutator, c int) error{}
	var lazySlice = []func(ctx context.Context, i *EntStepMutator, c int) error{}
	var postSlice = []func(ctx context.Context, i *ent.Step, c int) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.Step.Create()

	if b.recipeOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntStepMutator, c int) error {
			value := b.recipeOverride

			i.EntCreator().SetRecipe(value)

			i.Recipe = value
			return nil
		})
	} else {
		switch b.mutation.recipeType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.recipeFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.recipeFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.recipeFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.recipeFunc)
		}
	}

	if b.recipeIDOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntStepMutator, c int) error {
			value := b.recipeIDOverride

			i.EntCreator().SetRecipeID(value)

			i.RecipeID = value
			return nil
		})
	} else {
		switch b.mutation.recipeIDType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.recipeIDFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.recipeIDFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.recipeIDFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.recipeIDFunc)
		}
	}

	if b.textOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntStepMutator, c int) error {
			value := b.textOverride

			i.EntCreator().SetText(value)

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

	v := &EntStepMutator{}

	v._creator = entBuilder

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

	new, err := entBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

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
func (b *EntStepBuilder) CreateBatch(ctx context.Context, n int) ([]*ent.Step, error) {
	var results []*ent.Step
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *EntStepBuilder) CreateBatchV(ctx context.Context, n int) ([]ent.Step, error) {
	var results []ent.Step
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
