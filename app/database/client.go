package database

import (
	"database/sql"
	"fmt"
	
	_ "github.com/lib/pq"
)

func NewDatabaseConn(user string, password string, name string) (*sql.DB, error) {
	connArgs := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, name)
	conn, err := sql.Open("postgres", connArgs)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
