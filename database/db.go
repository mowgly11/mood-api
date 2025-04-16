package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Init() {
	db_cfg := mysql.NewConfig()
	db_cfg.User = "DB_USER"
	db_cfg.Passwd = "DB_PASSWORD"
	db_cfg.Net = "tcp"
	db_cfg.Addr = "127.0.0.1:3306"
	db_cfg.DBName = "mood"

	var err error
	DB, err = sql.Open("mysql", db_cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database")
}
