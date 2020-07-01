package recipe

import (
	"context"
	"database/sql"
	"fmt"

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

	fmt.Println("Ingredients: ", ingredients)

	// Remove ingredients that are already in DB
	for _, ing1 := range ingredients {
		for j, ing2 := range ingredientList {
			if ing1 == ing2 {
				ingredientList[j] = ingredientList[len(ingredientList)-1]
				ingredientList[len(ingredientList)-1] = ""
				ingredientList = ingredientList[:len(ingredientList)-1]
			}
		}
	}

	fmt.Println("IngredientList: ", ingredientList)

	// Then insert remaining ingredients into DB
	query := `INSERT INTO ` + constant.RR.Ingredients +
		` (name) VALUES`

	var args []interface{}
	for i, ing := range ingredientList {
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

	return nil
}
