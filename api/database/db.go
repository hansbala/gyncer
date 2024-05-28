package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var sqlConnection *sql.DB

func ConnectToDB() (*sql.DB, error) {
	// one (*sql.DB) is safe for concurrent use so return if connection is already created
	if sqlConnection != nil {
		return sqlConnection, nil
	}
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_ROOT_USER"),
		Passwd: os.Getenv("MYSQL_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	// Get a database handle.
	localSqlConnection, err := sql.Open("mysql", cfg.FormatDSN())
	// no need to defer closing DB connection since service should always be live
	if err != nil {
		return nil, err
	}
	sqlConnection = localSqlConnection

	// test ping the database
	err = sqlConnection.Ping()
	if err != nil {
		return nil, err
	}

	return sqlConnection, nil
}
