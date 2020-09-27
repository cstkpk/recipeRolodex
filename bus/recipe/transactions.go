package recipe

import (
	"context"
	"database/sql"
	"strings"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/logger"
)

// InsertRecipeDetails is the transaction that handles adding recipe details to the DB
func InsertRecipeDetails(ctx context.Context, tx *sql.Tx, season, title, author, link string) error {

	query := `INSERT INTO ` + constant.RR.Recipes +
		` (season, title, author, link) VALUES (?,?,?,?)`

	res, err := tx.ExecContext(ctx, query, season, title, author, link)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbInsertFailure
	}

	rows, err := res.RowsAffected()
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbQueryFailure
	}
	if rows != 1 {
		if rows == 0 {
			logger.Error.Println(logger.GetCallInfo(), constant.Errors.NoRowsAffected)
			return constant.Errors.NoRowsAffected
		}
		logger.Error.Println(logger.GetCallInfo(), constant.Errors.UnexpectedRowsAffected.Error())
		return constant.Errors.UnexpectedRowsAffected
	}

	logger.Info.Println(logger.GetCallInfo(), "Successfully inserted recipe into database")
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
		logger.Error.Println(logger.GetCallInfo(), err.Error())
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
			logger.Error.Println(logger.GetCallInfo(), err.Error())
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
		logger.Info.Println(logger.GetCallInfo(), "No new ingredients to insert")
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
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbInsertFailure
	}

	logger.Info.Printf("%s Successfully inserted %v ingredients", logger.GetCallInfo(), len(ingToAdd))
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
		logger.Error.Println(logger.GetCallInfo(), err.Error())
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
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var ingredientIDs []int64
	var ingredientID int64
	for rows.Next() {
		err = rows.Scan(&ingredientID)
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
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
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return constant.Errors.DbInsertFailure
		}
		rows, err := res.RowsAffected()
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return constant.Errors.DbQueryFailure
		}
		if rows != 1 {
			if rows == 0 {
				logger.Error.Println(logger.GetCallInfo(), constant.Errors.NoRowsAffected.Error())
				return constant.Errors.NoRowsAffected
			}
			logger.Error.Println(logger.GetCallInfo(), constant.Errors.UnexpectedRowsAffected.Error())
			return constant.Errors.UnexpectedRowsAffected
		}
	}

	logger.Info.Printf("%s Inserted link rows for ingredients with IDs %v for recipe with ID %v", logger.GetCallInfo(), ingredientIDs, recipeID)
	return nil
}
