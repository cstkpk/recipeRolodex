package busready

import (
	"context"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/logger"
	"github.com/cstkpk/recipeRolodex/mysql"
)

// GetReady pings the db to make sure that a connection can be established
func GetReady(ctx context.Context) error {
	// TODO: Create a connection pool to the db
	// For now just pinging it

	_, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbConnectionFailure
	}
	logger.Info.Println(logger.GetCallInfo(), "Successfully connected to the database")
	return nil
}
