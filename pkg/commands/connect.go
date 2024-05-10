package commands

import (
	"database/sql"
	"mysql-controller/pkg/env"

	"github.com/go-sql-driver/mysql"
)

func DBConnect() (*sql.DB, error) {

	host := env.Get("DB_HOST", "localhost")
	port := env.Get("DB_PORT", "3306")

	addr := host + ":" + port

	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   addr,
		DBName: "to-do-list",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}