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
