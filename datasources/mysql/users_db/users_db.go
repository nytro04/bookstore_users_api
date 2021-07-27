package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	UsersDB *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?", "root", "superM@N04", "127.0.0.1", "users_db")

	var err error
	UsersDB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful 🚀")
}
