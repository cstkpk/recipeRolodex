package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect initiates a connection to the database
func Connect(ctx context.Context, dbName string) (*sql.DB, error) {
	dbusername := os.Getenv("MARIA_USERNAME")
	dbpassword := os.Getenv("MARIA_PASSWORD")

	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}
	fmt.Println("Successfully connected to db")
	return db, nil
}
