package database

import (
	"database/sql"
)

func (d DatabaseClient) UserResgister(user string, email string, password string) (bool, error) {
	// comprobamos que el usuario no exista
	query, err := d.client.Prepare("SELECT true FROM users WHERE username = $1 OR email = $2 LIMIT 1")
	if err != nil {
		return false, err
	}
	defer query.Close()
	var exists bool
	if err := query.QueryRow(user, email).Scan(&exists); err != nil && err != sql.ErrNoRows {
		return false, err
	} else if exists {
		return false, nil
	}
	// si el usuario no existe lo registramos
	insert, err := d.client.Prepare("INSERT INTO users (username, email,password) VALUES ($1, $2, $3)")
	if err != nil {
		return false, err
	}
	defer insert.Close()
	_, err = insert.Exec(user, email, password)

	return true, err
}

func (d DatabaseClient) UserLogin(user string, password string) (string, error) {
	// obtenemos el usuario
	query, err := d.client.Prepare("SELECT id FROM users WHERE (username = $1 OR email = $1) AND password = $2 LIMIT 1")
	if err != nil {
		return "", err
	}
	defer query.Close()
	var id string
	if err := query.QueryRow(user, password).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
