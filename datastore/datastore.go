package datastore

import (
	"context"

	"github.com/jackc/pgx/v5"
)


type DataStore struct {
	
}


func NewDatastore() (string, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/demo_dev")
	if err != nil {
		return "", err
	}

	defer conn.Close(context.Background())

	var email string

	err = conn.QueryRow(context.Background(), "SELECT email FROM users WHERE id=$1", 1).Scan(&email)
	if err != nil {
		return email, err
	}

	return email, nil
	
}
