package constant

type databases struct {
	RecipeRolodex string
}

// DBs holds the databases used by this application
var DBs = databases{
	RecipeRolodex: "RecipeRolodex",
}

type rrTables struct {
	Ingredients string
	Link        string
	Recipes     string
}

// RR holds the tables within the RecipeRolodex DB
var RR = rrTables{
	Ingredients: DBs.RecipeRolodex + ".Ingredients",
	Link:        DBs.RecipeRolodex + ".Link",
	Recipes:     DBs.RecipeRolodex + ".Recipes",
}
