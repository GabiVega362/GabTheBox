package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	client *sql.DB
}

func NewDatabaseClient(user string, password string, name string) (*DatabaseClient, error) {
	connArgs := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, name)
	conn, err := sql.Open("postgres", connArgs)
	if err != nil {
		return nil, err
	}
	return &DatabaseClient{
		client: conn,
	}, nil
}

func (d DatabaseClient) UserResgister(email string, password string) (bool, error) {
	// comprobamos que el usuario no exista
	query, err := d.client.Prepare("SELECT true FROM users WHERE email = $1 LIMIT 1")
	if err != nil {
		return false, err
	}
	defer query.Close()
	var exists bool
	if err := query.QueryRow(email).Scan(&exists); err != nil && err != sql.ErrNoRows {
		return false, err
	} else if exists {
		return false, nil
	}
	// si el usuario no existe lo registramos
	insert, err := d.client.Prepare("INSERT INTO users (email,password) VALUES ($1, $2)")
	if err != nil {
		return false, err
	}
	defer insert.Close()
	_, err = insert.Exec(email, password)

	return true, err
}

func (d DatabaseClient) UserLogin(email string, password string) (uint, error) {
	// obtenemos el usuario
	query, err := d.client.Prepare("SELECT id FROM users WHERE email = $1 AND password = $2 LIMIT 1")
	if err != nil {
		return 0, err
	}
	defer query.Close()
	var id uint
	if err := query.QueryRow(email, password).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}