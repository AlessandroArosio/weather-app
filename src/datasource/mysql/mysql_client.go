package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mySqlUsername = "mysql_username"
	mySqlPassword = "mysql_password"
	mySqlHost = "mysql_host"
	mySqlSchema = "mysql_schema"
	driverName = "mysql"
)

var (
	username = os.Getenv(mySqlUsername)
	password = os.Getenv(mySqlPassword)
	host = os.Getenv(mySqlHost)
	schema = os.Getenv(mySqlSchema)
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
}
