package database

import (
	"fmt"
	"lms/models"

	"github.com/jmoiron/sqlx"
)

func Login(db sqlx.Ext, username string, password string) (*models.LoginResponsedb, error) {

	counter := 0
	var id int64
	var user string
	var pass string

	rows, err := db.Queryx(`SELECT ID, Username, Password FROM Admin
	WHERE Username=? AND Password=?`,
		username, password)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &user, &pass)
		if err != nil {
			return nil, err
		}
		counter++
	}

	if counter == 0 {
		return nil, fmt.Errorf("User does not exists")
	}

	return &models.LoginResponsedb{
		ID:       id,
		Username: user,
		Password: pass,
	}, nil
}
