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
	Client   *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)

	// username = "root"
	// password = "root"
	// host     = "103.186.30.178:35200"
	// schema   = "udemy_users_db"

	// username = "root"
	// password = "root"
	// host     = "localhost:3306"
	// schema   = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	// gaboleh naro log kaya gini karena bahaya
	// log.Println(fmt.Sprintf("about to connect tot %s", dataSourceName))
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	// mysql.SetLogger()

	log.Println("database successFuly configured")

}
