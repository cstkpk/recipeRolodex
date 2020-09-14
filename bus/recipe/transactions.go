package recipe

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cstkpk/recipeRolodex/constant"
)

// InsertRecipeDetails is the transaction that handles adding recipe details to the DB
func InsertRecipeDetails(ctx context.Context, tx *sql.Tx, season, title, author, link string) error {

	query := `INSERT INTO ` + constant.RR.Recipes +
		` (season, title, author, link) VALUES (?,?,?,?)`

	res, err := tx.ExecContext(ctx, query, season, title, author, link)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return constant.Errors.DbInsertFailure
	}

	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return constant.Errors.DbQueryFailure
	}
	if rows != 1 {
		if rows == 0 {
			fmt.Println("Error:", constant.Errors.NoRowsAffected)
			return constant.Errors.NoRowsAffected
		}
		fmt.Println("Error:", constant.Errors.UnexpectedRowsAffected)
		return constant.Errors.UnexpectedRowsAffected
	}

	fmt.Println("Info: Successfully inserted recipe into database")
	return nil
}

// InsertIngredients checks the Ingredients table to make sure the recipe's ingredients exist
// and if not adds them to the table
func InsertIngredients(ctx context.Context, tx *sql.Tx, ingredientList []string) error {
	// First make sure all ingredients are lowercase
	for i, ing := range ingredientList {
		ingredientList[i] = strings.ToLower(ing)
	}

	// First check to see which ingredients are already in the Ingredients table
	query1 := `SELECT name FROM ` + constant.RR.Ingredients +
		` WHERE name = ?`

	var args1 []interface{}
	for i, ing := range ingredientList {
		args1 = append(args1, ing)
		if i != 0 {
			query1 += " OR name = ?"
		}
	}

	rows, err := tx.QueryContext(ctx, query1, args1...)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var ingredients []string
	for rows.Next() {
		var ingredient string
		err = rows.Scan(
			&ingredient,
		)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return constant.Errors.InternalServer
		}
		ingredients = append(ingredients, ingredient)
	}

	// Remove ingredients that are already in DB
	var ingToAdd []string
	copy(ingToAdd, ingredientList)
	for _, ing1 := range ingredients {
		for j, ing2 := range ingToAdd {
			if ing1 == ing2 {
				ingToAdd[j] = ingToAdd[len(ingToAdd)-1]
				ingToAdd[len(ingToAdd)-1] = ""
				ingToAdd = ingToAdd[:len(ingToAdd)-1]
			}
		}
	}

	// If there are no new ingredients to insert, return
	if len(ingToAdd) == 0 {
		fmt.Println("Info: no new ingredients to insert")
		return nil
	}

	// Then insert remaining ingredients into DB
	query := `INSERT INTO ` + constant.RR.Ingredients +
		` (name) VALUES`

	var args []interface{}
	for i, ing := range ingToAdd {
		args = append(args, ing)
		if i == 0 {
			query += " (?)"
		} else {
			query += ", (?)"
		}
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return constant.Errors.DbInsertFailure
	}

	fmt.Printf("Info: Successfully inserted %v ingredients\n", len(ingToAdd))
	return nil
}

// InsertLink creates the Link row(s) in the DB, which links together the
// recipe and its ingredients
func InsertLink(ctx context.Context, tx *sql.Tx, recipeTitle string, ingredients []string) error {
	// First find recipe ID
	queryRecipe := `SELECT id FROM ` + constant.RR.Recipes +
		` WHERE title = ?`
	var recipeID int64
	err := tx.QueryRowContext(ctx, queryRecipe, recipeTitle).Scan(&recipeID)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return constant.Errors.DbQueryFailure
	}

	// Then find all of the ingredient IDs
	queryIngredients := `SELECT id FROM ` + constant.RR.Ingredients +
		` WHERE name = ?`
	var args []interface{}
	for i, ing := range ingredients {
		args = append(args, ing)
		if i != 0 {
			queryIngredients += " OR name = ?"
		}
	}
	rows, err := tx.QueryContext(ctx, queryIngredients, args...)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var ingredientIDs []int64
	var ingredientID int64
	for rows.Next() {
		err = rows.Scan(&ingredientID)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return constant.Errors.DbQueryFailure
		}
		ingredientIDs = append(ingredientIDs, ingredientID)
	}

	// Last, insert a row for each ingredient ID corresponding to the recipe ID
	queryLink := `INSERT INTO ` + constant.RR.Link +
		` (recipeID, ingredientID) VALUES (?,?)`

	for _, ing := range ingredientIDs {
		res, err := tx.ExecContext(ctx, queryLink, recipeID, ing)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return constant.Errors.DbInsertFailure
		}
		rows, err := res.RowsAffected()
		if err != nil {
			fmt.Println("Error:", err.Error())
			return constant.Errors.DbQueryFailure
		}
		if rows != 1 {
			if rows == 0 {
				fmt.Println("Error:", constant.Errors.NoRowsAffected.Error())
				return constant.Errors.NoRowsAffected
			}
			fmt.Println("Error:", constant.Errors.UnexpectedRowsAffected.Error())
			return constant.Errors.UnexpectedRowsAffected
		}
	}

	fmt.Printf("Info: Inserted link rows for ingredients with IDs %v for recipe with ID %v \n", ingredientIDs, recipeID)
	return nil
}
