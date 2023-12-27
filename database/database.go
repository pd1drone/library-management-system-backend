package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitializeLMSDatabase(dbname, username, password, dbhost, dbport string) (*sqlx.DB, error) {
	conn := username + ":" + password + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	lmsdb, err := sqlx.Connect("mysql", conn)

	if err != nil {
		return nil, fmt.Errorf("Error in initializing watchtower database: %s", err)
	}

	return lmsdb, nil
}
