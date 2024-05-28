package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var sqlConnection *sql.DB

// this should only be called once when the service first receives a request
func instantiateDBConnection() error {
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
		return err
	}
	sqlConnection = localSqlConnection

	// test ping the database
	err = sqlConnection.Ping()
	if err != nil {
		return err
	}
	return nil
}

// if a db connection doesn't exist yet, create one and send it over. if not, use the pre-exisiting db connection
func ConnectToDB() (*sql.DB, error) {
	if sqlConnection == nil {
		if err := instantiateDBConnection(); err != nil {
			return nil, err
		}
	}
	// one (*sql.DB) is safe for concurrent use so return if connection is already created
	return sqlConnection, nil
}
