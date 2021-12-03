package model

type Recipe struct {
	Name        string
	AuthorName  string
	Servings    int
	Steps       []*RecipeStep
	Ingredients []*RecipeIngredient
	Tags        []*Category
	Related     []*Recipe
}

type RecipeStep struct {
	Text string
}

type RecipeIngredient struct {
	Ingredient *Ingredient
	Quantity   float32
	Unit       string
}

type Ingredient struct {
	Name string
	Tags []*Category
}

type Category struct {
	Name string
}

type User struct {
	Name    string
	Tags    []*Category
	Recipes []*Recipe
}
