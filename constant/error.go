package constant

import "github.com/cstkpk/recipeRolodex/rrerror"

// Errors contains RR-specific errors to return to user
var Errors = struct {
	// DB errors
	DbConnectionFailure error
	DbQueryFailure      error
	InternalService     error

	// Endpoint errors
	NoRecipesFound error
}{
	// DB errors
	DbConnectionFailure: rrerror.New("There was an error connecting to the database", 500),
	DbQueryFailure:      rrerror.New("There was an error querying the database", 500),
	InternalService:     rrerror.New("Internal service error", 500),

	// Endpoint errors
	NoRecipesFound: rrerror.New("No recipes found that match search criteria", 404),
}
