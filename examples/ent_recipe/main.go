package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yiling-J/carrier/examples/ent_recipe/carrier"
	"github.com/Yiling-J/carrier/examples/ent_recipe/carrier/factory"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/category"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/recipe"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/user"
	_ "github.com/mattn/go-sqlite3"
)

func initFactory() *carrier.EntFactory {
	client, err := ent.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	factory := carrier.NewEntFactory(client)
	initCategoryFactory(factory)
	initIngredientFactory(factory)
	initRecipeFactory(factory)
	initRecipeIngredientFactory(factory)
	initStepFactory(factory)
	initUserFactory(factory)
	return factory
}

func initRecipeFactory(f *carrier.EntFactory) {
	meta := carrier.EntRecipeMetaFactory().
		SetNameSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("recipe_%d", i), nil
		}).
		SetServingsDefault(3).
		SetAuthorFactory(func(ctx context.Context) (*ent.User, error) {
			return f.UserFactory().Create(ctx)
		}).
		SetStepsCountPostFunc(
			func(ctx context.Context, set bool, obj *ent.Recipe, i int) error {
				count := 3
				if set {
					count = i
				}
				if count > 0 {
					steps, err := f.StepFactory().CreateBatch(ctx, count)
					if err != nil {
						return err
					}
					_, err = obj.Update().AddSteps(steps...).Save(ctx)
					return err
				}
				return nil
			},
		).
		SetIngredientsCountPostFunc(
			func(ctx context.Context, set bool, obj *ent.Recipe, i int) error {
				count := 6
				if set {
					count = i
				}
				if count > 0 {
					ingredients, err := f.RecipeIngredientFactory().CreateBatch(ctx, count)
					if err != nil {
						return err
					}
					_, err = obj.Update().AddRecipeIngredients(ingredients...).Save(ctx)
					return err
				}
				return nil
			},
		).
		SetVeganTrait(
			factory.EntRecipeTrait().SetAfterCreateFunc(
				func(ctx context.Context, i *ent.Recipe) error {
					c, err := f.Client().Category.Query().
						Where(category.Name("vegan")).First(ctx)
					if err != nil {
						c, err = f.CategoryFactory().SetName("vegan").Create(ctx)
						if err != nil {
							return err
						}
					}
					_, err = i.Update().AddTags(c).Save(ctx)
					return err
				},
			),
		)
	f.SetRecipeFactory(meta.Build())
}

func initIngredientFactory(f *carrier.EntFactory) {
	meta := carrier.EntIngredientMetaFactory().
		SetNameSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("ingredient_%d", i), nil
		})
	f.SetIngredientFactory(meta.Build())
}

func initUserFactory(f *carrier.EntFactory) {
	meta := carrier.EntUserMetaFactory().
		SetNameSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("user_%d", i), nil
		}).
		SetRecipesPostFunc(func(ctx context.Context, set bool, obj *ent.User, i int) error {
			if set && i > 0 {
				_, err := f.RecipeFactory().SetAuthor(obj).CreateBatch(ctx, i)
				if err != nil {
					return err
				}
				return nil
			}
			return nil
		})
	f.SetUserFactory(meta.Build())
}

func initCategoryFactory(f *carrier.EntFactory) {
	meta := carrier.EntCategoryMetaFactory().
		SetNameDefault("cat")
	f.SetCategoryFactory(meta.Build())
}

func initStepFactory(f *carrier.EntFactory) {
	meta := carrier.EntStepMetaFactory().
		SetTextSequence(func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("step_%d", i), nil
		})
	f.SetStepFactory(meta.Build())
}

func initRecipeIngredientFactory(f *carrier.EntFactory) {
	meta := carrier.EntRecipeIngredientMetaFactory().
		SetIngredientFactory(func(ctx context.Context) (*ent.Ingredient, error) {
			return f.IngredientFactory().Create(ctx)
		}).
		SetQuantityDefault(4.5).
		SetUnitDefault("gram")
	f.SetRecipeIngredientFactory(meta.Build())
}

func main() {
	factory := initFactory()
	// user with 3 recipes, result in user.json
	u, err := factory.UserFactory().SetRecipesPost(3).Create(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(u.ID)
	u, err = factory.Client().User.Query().Where(user.IDEQ(u.ID)).WithRecipes(func(rq *ent.RecipeQuery) {
		rq.WithAuthor().WithSteps().WithRecipeIngredients(func(riq *ent.RecipeIngredientQuery) {
			riq.WithIngredient()
		}).WithSteps().WithTags()
	}).First(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// vegan recipe with 5 steps and 15 ingredients, result in recipe.json
	factory = initFactory()
	r, err := factory.RecipeFactory().
		WithVeganTrait().
		SetStepsCountPost(5).
		SetIngredientsCountPost(15).Create(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r, err = factory.Client().Recipe.Query().Where(recipe.IDEQ(r.ID)).
		WithAuthor().WithSteps().WithRecipeIngredients(func(riq *ent.RecipeIngredientQuery) {
		riq.WithIngredient()
	}).WithSteps().WithTags().First(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
