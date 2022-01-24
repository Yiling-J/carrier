// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/examples/recipe/model"

	"context"
)

type ingredientMutation struct {
	nameType int
	nameFunc func(ctx context.Context, i *model.Ingredient, c int) error

	tagsType int
	tagsFunc func(ctx context.Context, i *model.Ingredient, c int) error

	afterCreateFunc func(ctx context.Context, i *model.Ingredient) error
}
type IngredientMetaFactory struct {
	mutation ingredientMutation
}
type ingredientTrait struct {
	mutation ingredientMutation
	updates  []func(m *ingredientMutation)
}

func IngredientTrait() *ingredientTrait {
	return &ingredientTrait{}
}
func (*ingredientMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *model.Ingredient) error) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.afterCreateFunc = fn
	}
}

func (*ingredientMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *model.Ingredient, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Name = value
			return nil
		}
	}
}
func (*ingredientMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *model.Ingredient) (string, error)) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *model.Ingredient, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Name = value
			return nil
		}
	}
}
func (*ingredientMutation) nameDefaultMutateFunc(v string) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *model.Ingredient, c int) error {

			i.Name = v
			return nil
		}
	}
}
func (*ingredientMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *model.Ingredient, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Name = value

			return nil
		}
	}
}

func (f *IngredientMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *IngredientMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *IngredientMetaFactory) SetNameLazy(fn func(ctx context.Context, i *model.Ingredient) (string, error)) *IngredientMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *IngredientMetaFactory) SetNameDefault(v string) *IngredientMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *IngredientMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *IngredientMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *ingredientTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}
func (t *ingredientTrait) SetNameLazy(fn func(ctx context.Context, i *model.Ingredient) (string, error)) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}
func (t *ingredientTrait) SetNameDefault(v string) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}
func (t *ingredientTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

func (*ingredientMutation) tagsSequenceMutateFunc(fn func(ctx context.Context, i int) ([]*model.Category, error)) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.tagsType = TypeSequence
		m.tagsFunc = func(ctx context.Context, i *model.Ingredient, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Tags = value
			return nil
		}
	}
}
func (*ingredientMutation) tagsLazyMutateFunc(fn func(ctx context.Context, i *model.Ingredient) ([]*model.Category, error)) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.tagsType = TypeLazy
		m.tagsFunc = func(ctx context.Context, i *model.Ingredient, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Tags = value
			return nil
		}
	}
}
func (*ingredientMutation) tagsDefaultMutateFunc(v []*model.Category) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.tagsType = TypeDefault
		m.tagsFunc = func(ctx context.Context, i *model.Ingredient, c int) error {

			i.Tags = v
			return nil
		}
	}
}
func (*ingredientMutation) tagsFactoryMutateFunc(fn func(ctx context.Context) ([]*model.Category, error)) func(m *ingredientMutation) {
	return func(m *ingredientMutation) {
		m.tagsType = TypeFactory
		m.tagsFunc = func(ctx context.Context, i *model.Ingredient, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Tags = value

			return nil
		}
	}
}

func (f *IngredientMetaFactory) SetTagsSequence(fn func(ctx context.Context, i int) ([]*model.Category, error)) *IngredientMetaFactory {
	f.mutation.tagsSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *IngredientMetaFactory) SetTagsLazy(fn func(ctx context.Context, i *model.Ingredient) ([]*model.Category, error)) *IngredientMetaFactory {
	f.mutation.tagsLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *IngredientMetaFactory) SetTagsDefault(v []*model.Category) *IngredientMetaFactory {
	f.mutation.tagsDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *IngredientMetaFactory) SetTagsFactory(fn func(ctx context.Context) ([]*model.Category, error)) *IngredientMetaFactory {
	f.mutation.tagsFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *ingredientTrait) SetTagsSequence(fn func(ctx context.Context, i int) ([]*model.Category, error)) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.tagsSequenceMutateFunc(fn))
	return t
}
func (t *ingredientTrait) SetTagsLazy(fn func(ctx context.Context, i *model.Ingredient) ([]*model.Category, error)) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.tagsLazyMutateFunc(fn))
	return t
}
func (t *ingredientTrait) SetTagsDefault(v []*model.Category) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.tagsDefaultMutateFunc(v))
	return t
}
func (t *ingredientTrait) SetTagsFactory(fn func(ctx context.Context) ([]*model.Category, error)) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.tagsFactoryMutateFunc(fn))
	return t
}

func (f *IngredientMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *model.Ingredient) error) *IngredientMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}
func (t *ingredientTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *model.Ingredient) error) *ingredientTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

func (f *IngredientMetaFactory) Build() *IngredientFactory {
	return &IngredientFactory{meta: *f, counter: &Counter{}}
}

type IngredientFactory struct {
	meta    IngredientMetaFactory
	counter *Counter
}

func (f *IngredientFactory) SetName(i string) *IngredientBuilder {
	builder := &IngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)

	return builder
}

func (f *IngredientFactory) SetTags(i []*model.Category) *IngredientBuilder {
	builder := &IngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetTags(i)

	return builder
}

func (f *IngredientFactory) Create(ctx context.Context) (*model.Ingredient, error) {
	builder := &IngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.Create(ctx)
}
func (f *IngredientFactory) CreateV(ctx context.Context) (model.Ingredient, error) {
	builder := &IngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateV(ctx)
}
func (f *IngredientFactory) CreateBatch(ctx context.Context, n int) ([]*model.Ingredient, error) {
	builder := &IngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatch(ctx, n)
}
func (f *IngredientFactory) CreateBatchV(ctx context.Context, n int) ([]model.Ingredient, error) {
	builder := &IngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatchV(ctx, n)
}

type IngredientBuilder struct {
	factory  *IngredientFactory
	mutation ingredientMutation
	counter  *Counter

	nameOverride  string
	nameOverriden bool

	tagsOverride  []*model.Category
	tagsOverriden bool
}

func (b *IngredientBuilder) SetName(i string) *IngredientBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

func (b *IngredientBuilder) SetTags(i []*model.Category) *IngredientBuilder {
	b.tagsOverride = i
	b.tagsOverriden = true
	return b
}

func (b *IngredientBuilder) CreateV(ctx context.Context) (model.Ingredient, error) {
	var d model.Ingredient
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

func (b *IngredientBuilder) Create(ctx context.Context) (*model.Ingredient, error) {

	var preSlice = []func(ctx context.Context, i *model.Ingredient, c int) error{}
	var lazySlice = []func(ctx context.Context, i *model.Ingredient, c int) error{}
	var postSlice = []func(ctx context.Context, i *model.Ingredient, c int) error{}

	index := b.counter.Get()
	_ = index

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.Ingredient, c int) error {
			value := b.nameOverride

			i.Name = value
			return nil
		})
	} else {
		switch b.mutation.nameType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.nameFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.nameFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.nameFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.nameFunc)
		}
	}

	if b.tagsOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.Ingredient, c int) error {
			value := b.tagsOverride

			i.Tags = value
			return nil
		})
	} else {
		switch b.mutation.tagsType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.tagsFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.tagsFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.tagsFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.tagsFunc)
		}
	}

	v := &model.Ingredient{}
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
func (b *IngredientBuilder) CreateBatch(ctx context.Context, n int) ([]*model.Ingredient, error) {
	var results []*model.Ingredient
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *IngredientBuilder) CreateBatchV(ctx context.Context, n int) ([]model.Ingredient, error) {
	var results []model.Ingredient
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}