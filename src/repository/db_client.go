package repository

import (
	"database/sql"
	"fmt"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/config"
	"github.com/JaimePalomo/nfcliftserver-ddd/toolkit/log"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLClient() *sql.DB {
	var err error
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.Cfg.Mysql_nfclift_username,
		config.Cfg.Mysql_nfclift_password,
		config.Cfg.Mysql_nfclift_host,
		config.Cfg.Mysql_nfclift_schema,
	)
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Info("Database succesfully configured and connected!")

	return Client
}
