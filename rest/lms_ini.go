package rest

import (
	"fmt"
	"lms/database"

	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

type LMSConfig struct {
	LMSdb *sqlx.DB
}

func New() (*LMSConfig, error) {

	// read config file
	cfg, err := ini.Load("/root/library-management-system-backend/config.ini")
	if err != nil {
		return nil, fmt.Errorf("Fail to read file: %v", err)
	}

	dbSection := cfg.Section("db")
	user := dbSection.Key("user").String()
	password := dbSection.Key("password").String()
	dbhost := dbSection.Key("dbhost").String()
	dbport := dbSection.Key("dbport").String()
	dbname := dbSection.Key("dbname").String()

	lmsdb, err := database.InitializeLMSDatabase(dbname, user, password, dbhost, dbport)
	if err != nil {
		return nil, err
	}

	return &LMSConfig{
		lmsdb,
	}, nil
}
