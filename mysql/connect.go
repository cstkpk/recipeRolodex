package mysql

import (
	"context"
	"database/sql"
	"os"

	"github.com/cstkpk/recipeRolodex/logger"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect initiates a connection to the database
func Connect(ctx context.Context, dbName string) (*sql.DB, error) {
	dbusername := os.Getenv("MARIA_USERNAME")
	dbpassword := os.Getenv("MARIA_PASSWORD")

	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error.Println("Error:", err.Error())
		return nil, err
	}

	return db, nil
}
