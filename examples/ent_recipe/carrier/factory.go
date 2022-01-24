// Code generated by carrier, DO NOT EDIT.
package carrier

import (
	"github.com/Yiling-J/carrier/examples/ent_recipe/carrier/factory"

	"github.com/Yiling-J/carrier/examples/ent_recipe/ent"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

type EntFactory struct {
	recipeFactory *factory.EntRecipeFactory

	userFactory *factory.EntUserFactory

	ingredientFactory *factory.EntIngredientFactory

	categoryFactory *factory.EntCategoryFactory

	recipeIngredientFactory *factory.EntRecipeIngredientFactory

	stepFactory *factory.EntStepFactory

	client *ent.Client
}

func NewEntFactory(client *ent.Client) *EntFactory {
	return &EntFactory{client: client}
}

func (f *EntFactory) Client() *ent.Client {
	return f.client
}

func EntRecipeMetaFactory() *factory.EntRecipeMetaFactory {
	return &factory.EntRecipeMetaFactory{}
}
func (f *EntFactory) SetRecipeFactory(c *factory.EntRecipeFactory) *EntFactory {
	f.recipeFactory = c.Client(f.client)
	return f
}
func (f *EntFactory) RecipeFactory() *factory.EntRecipeFactory {
	return f.recipeFactory
}

func EntUserMetaFactory() *factory.EntUserMetaFactory {
	return &factory.EntUserMetaFactory{}
}
func (f *EntFactory) SetUserFactory(c *factory.EntUserFactory) *EntFactory {
	f.userFactory = c.Client(f.client)
	return f
}
func (f *EntFactory) UserFactory() *factory.EntUserFactory {
	return f.userFactory
}

func EntIngredientMetaFactory() *factory.EntIngredientMetaFactory {
	return &factory.EntIngredientMetaFactory{}
}
func (f *EntFactory) SetIngredientFactory(c *factory.EntIngredientFactory) *EntFactory {
	f.ingredientFactory = c.Client(f.client)
	return f
}
func (f *EntFactory) IngredientFactory() *factory.EntIngredientFactory {
	return f.ingredientFactory
}

func EntCategoryMetaFactory() *factory.EntCategoryMetaFactory {
	return &factory.EntCategoryMetaFactory{}
}
func (f *EntFactory) SetCategoryFactory(c *factory.EntCategoryFactory) *EntFactory {
	f.categoryFactory = c.Client(f.client)
	return f
}
func (f *EntFactory) CategoryFactory() *factory.EntCategoryFactory {
	return f.categoryFactory
}

func EntRecipeIngredientMetaFactory() *factory.EntRecipeIngredientMetaFactory {
	return &factory.EntRecipeIngredientMetaFactory{}
}
func (f *EntFactory) SetRecipeIngredientFactory(c *factory.EntRecipeIngredientFactory) *EntFactory {
	f.recipeIngredientFactory = c.Client(f.client)
	return f
}
func (f *EntFactory) RecipeIngredientFactory() *factory.EntRecipeIngredientFactory {
	return f.recipeIngredientFactory
}

func EntStepMetaFactory() *factory.EntStepMetaFactory {
	return &factory.EntStepMetaFactory{}
}
func (f *EntFactory) SetStepFactory(c *factory.EntStepFactory) *EntFactory {
	f.stepFactory = c.Client(f.client)
	return f
}
func (f *EntFactory) StepFactory() *factory.EntStepFactory {
	return f.stepFactory
}