package busready

import (
	"context"
	"fmt"

	"github.com/cstkpk/recipeRolodex/mysql"
)

// GetReady pings the db to make sure that a connection can be established
func GetReady(ctx context.Context) error {
	// TODO: Create a connection pool to the db
	// For now just pinging it

	_, err := mysql.Connect(ctx)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}
	return nil
}
