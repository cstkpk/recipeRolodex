package constant

import "github.com/cstkpk/recipeRolodex/rrerror"

// Errors contains RR-specific errors to return to user
var Errors = struct {
	// DB errors
	DbConnectionFailure    error
	DbQueryFailure         error
	DbInsertFailure        error
	InternalServer         error
	NoRowsAffected         error
	UnexpectedRowsAffected error

	// Endpoint errors
	NoRecipesFound  error
	NoRecipeIDFound error
}{
	// DB errors
	DbConnectionFailure:    rrerror.New("There was an error connecting to the database", 500),
	DbQueryFailure:         rrerror.New("There was an error querying the database", 500),
	DbInsertFailure:        rrerror.New("There was an error adding your request to the database", 500),
	InternalServer:         rrerror.New("Internal server error", 500),
	NoRowsAffected:         rrerror.New("No changes made", 400),
	UnexpectedRowsAffected: rrerror.New("Unexpected number of rows affected", 400),

	// Endpoint errors
	NoRecipesFound:  rrerror.New("No recipes found that match search criteria", 404),
	NoRecipeIDFound: rrerror.New("No recipe found with the specified ID", 404),
}
