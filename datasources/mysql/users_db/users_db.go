package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nytro04/bookstore_users_api/utils/env_utils"
)

// const (
// 	mysql_users_username = "mysql_users_username"
// 	mysql_users_password = "mysql_users_password"
// 	mysql_users_host     = "mysql_users_host"
// 	mysql_users_schema   = "mysql_users_schema"
// )

var (
	UsersDB *sql.DB
	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// schema   = os.Getenv(mysql_users_schema)
)

func init() {
	config, configErr := env_utils.LoadConfig(".")
	if configErr != nil {
		log.Fatal("Cannot load config:", configErr)
	}
 
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?", config.SQLUsername, config.SQLPassword, config.SQLHost, config.SQLSchema)
	//dataSource  eName = dataSourceName + "allowNativePasswords=false"
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
