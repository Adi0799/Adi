package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMySQL() {
	var err error
	connStr := "root:Player@079@tcp(127.0.0.1:3306)/notif_service"
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot reach MySQL:", err)
	}

	log.Println("Connected to MySQL database!")
}
