package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_ROOT_USER"),
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	// TODO: remove
	// test ping the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
