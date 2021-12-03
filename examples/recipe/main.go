package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yiling-J/carrier/examples/recipe/carrier"
	"github.com/Yiling-J/carrier/examples/recipe/carrier/factory"
	"github.com/Yiling-J/carrier/examples/recipe/model"
)

func initFactory() *carrier.Factory {
	factory := carrier.NewFactory()
	initRecipeFactory(factory)
	initIngredientFactory(factory)
	initUserFactory(factory)
	initStepFactory(factory)
	initCategoryFactory(factory)
	initRecipeIngredientFactory(factory)
	return factory
}

func initRecipeFactory(f *carrier.Factory) {
	meta := carrier.RecipeMetaFactory().
		SetNameSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("recipe_%d", i), nil
		}).
		SetServingsDefault(3).
		SetAuthorNameFactory(func(ctx context.Context) (string, error) {
			a, err := f.UserFactory().Create(ctx)
			if err != nil {
				return "", err
			}
			return a.Name, nil
		}).
		SetStepsCountPostFunc(
			func(ctx context.Context, set bool, obj *model.Recipe, i int) error {
				count := 3
				if set {
					count = i
				}
				if count > 0 {
					steps, err := f.RecipeStepFactory().CreateBatch(ctx, count)
					if err != nil {
						return err
					}
					obj.Steps = steps
				}
				return nil
			},
		).
		SetIngredientsCountPostFunc(
			func(ctx context.Context, set bool, obj *model.Recipe, i int) error {
				count := 6
				if set {
					count = i
				}
				if count > 0 {
					ingredients, err := f.RecipeIngredientFactory().CreateBatch(ctx, count)
					if err != nil {
						return err
					}
					obj.Ingredients = ingredients
				}
				return nil
			},
		).
		SetVeganTrait(
			factory.RecipeTrait().SetTagsFactory(
				func(ctx context.Context) ([]*model.Category, error) {
					return f.CategoryFactory().SetName("vegan").CreateBatch(ctx, 1)
				},
			),
		)
	f.SetRecipeFactory(meta.Build())
}

func initIngredientFactory(f *carrier.Factory) {
	meta := carrier.IngredientMetaFactory().
		SetNameSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("ingredient_%d", i), nil
		})
	f.SetIngredientFactory(meta.Build())
}

func initUserFactory(f *carrier.Factory) {
	meta := carrier.UserMetaFactory().
		SetNameSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("user_%d", i), nil
		}).
		SetRecipesPostFunc(func(ctx context.Context, set bool, obj *model.User, i int) error {
			if set && i > 0 {
				recipes, err := f.RecipeFactory().SetAuthorName(obj.Name).CreateBatch(ctx, i)
				if err != nil {
					return err
				}
				obj.Recipes = recipes
			}
			return nil
		})
	f.SetUserFactory(meta.Build())
}

func initCategoryFactory(f *carrier.Factory) {
	meta := carrier.CategoryMetaFactory().
		SetNameDefault("cat")
	f.SetCategoryFactory(meta.Build())
}

func initStepFactory(f *carrier.Factory) {
	meta := carrier.RecipeStepMetaFactory().
		SetTextSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("step_%d", i), nil
		})
	f.SetRecipeStepFactory(meta.Build())
}

func initRecipeIngredientFactory(f *carrier.Factory) {
	meta := carrier.RecipeIngredientMetaFactory().
		SetIngredientFactory(func(ctx context.Context) (*model.Ingredient, error) {
			return f.IngredientFactory().Create(ctx)
		}).
		SetQuantityDefault(4.5).
		SetUnitDefault("gram")
	f.SetRecipeIngredientFactory(meta.Build())
}

func main() {
	factory := initFactory()
	// user with 3 recipes, result in user.json
	user, err := factory.UserFactory().SetRecipesPost(3).Create(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// vegan recipe with 5 steps and 15 ingredients, result in recipe.json
	factory = initFactory()
	recipe, err := factory.RecipeFactory().
		WithVeganTrait().
		SetStepsCountPost(5).
		SetIngredientsCountPost(15).Create(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = json.Marshal(recipe)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
