package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	UsersDB  *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?", username, password, host, schema)
	//dataSourceName = dataSourceName + "allowNativePasswords=false"
	var err error
	UsersDB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = UsersDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful 🚀")
}
