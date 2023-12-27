package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type LoginResponse struct {
	ID       int64  `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func Login(db sqlx.Ext, username string, password string) (*LoginResponse, error) {

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

	return &LoginResponse{
		ID:       id,
		Username: user,
		Password: pass,
	}, nil
}
