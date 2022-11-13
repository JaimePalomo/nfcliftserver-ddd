package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysql_nfclift_username = "mysql_users_username"
	mysql_nfclift_password = "mysql_users_password"
	mysql_nfclift_host     = "mysql_users_host"
	mysql_nfclift_schema   = "mysql_users_schema"
)

func New() *sql.DB {
	username := os.Getenv(mysql_nfclift_username)
	password := os.Getenv(mysql_nfclift_password)
	host := os.Getenv(mysql_nfclift_host)
	schema := os.Getenv(mysql_nfclift_schema)
	var err error
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database succesfully configured and connected!")

	return Client
}
