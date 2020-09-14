package shop_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/mysql"
	"github.com/gavv/httpexpect"
)

// Global variables
var tdb *sql.DB

var (
	rrHost = "http://0.0.0.0:"
	rrPort = "6010"
)

func TestMain(m *testing.M) {
	ctx := context.Context(context.Background())
	db, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	tdb = db

	// Run tests
	v := m.Run()
	// Exit program
	os.Exit(v)
}

func TestReady(t *testing.T) {
	tt := []struct {
		name       string
		path       string
		statusCode int
	}{
		{"api is ready", "/rolodex/ready", 200},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  rrHost + rrPort,
				Reporter: httpexpect.NewAssertReporter(t),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(t, true),
				},
			})
			e.GET(tc.path).
				Expect().
				Status(tc.statusCode)
		})
	}
}

func TestRecipe(t *testing.T) {
	tt := []struct {
		name       string
		path       string
		id         int64
		statusCode int
	}{
		{"get recipe details", "/rolodex/recipe", 1, 200},
		{"invalid recipe id", "/rolodex/recipe", 100000, 404},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  rrHost + rrPort,
				Reporter: httpexpect.NewAssertReporter(t),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(t, true),
				},
			})
			e.GET(tc.path).
				WithQuery("recipeID", tc.id).
				Expect().
				Status(tc.statusCode)
		})
	}

	tt2 := []struct {
		name           string
		path           string
		author         string
		ingredientList string
		link           string
		season         string
		title          string
		statusCode     int
	}{
		{"add new recipe", "/rolodex/recipe", "Test Author", "carrots", "fakelink", "Winter", "Test Title", 200},
	}

	for _, tc := range tt2 {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  rrHost + rrPort,
				Reporter: httpexpect.NewAssertReporter(t),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(t, true),
				},
			})
			e.POST(tc.path).
				WithJSON(models.NewRecipe{Author: &tc.author, IngredientList: []string{tc.ingredientList}, Link: &tc.link, Season: &tc.season, Title: &tc.title}).
				Expect().
				Status(tc.statusCode)
		})
	}

	row := tdb.QueryRow("SELECT id FROM " + constant.RR.Recipes + " WHERE title = \"Test Title\"")
	var id int64
	err := row.Scan(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, err = tdb.Exec(fmt.Sprintf("DELETE FROM Link WHERE recipeID = %d", id))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, err = tdb.Exec(fmt.Sprintf("DELETE FROM recipes WHERE id = %d", id))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

}

func TestRecipes(t *testing.T) {
	tt := []struct {
		name        string
		path        string
		season      string
		ingredient1 string
		statusCode  int
	}{
		{"get all recipes", "/rolodex/recipes", "", "", 200},
		{"get recipes by season", "/rolodex/recipes", "winter", "", 200},
		{"get recipes by ingredient", "/rolodex/recipes", "", "avocado", 200},
		{"find no recipes by ingredient", "/rolodex/recipes", "", "fake_ingredient", 404},
	}

	type RecipesQueryObj struct {
		Season      string `url:"season"`
		Ingredient1 string `url:"ingredient1"`
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  rrHost + rrPort,
				Reporter: httpexpect.NewAssertReporter(t),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(t, true),
				},
			})
			e.GET(tc.path).
				WithQueryObject(RecipesQueryObj{Season: tc.season, Ingredient1: tc.ingredient1}).
				Expect().
				Status(tc.statusCode)
		})
	}
}
