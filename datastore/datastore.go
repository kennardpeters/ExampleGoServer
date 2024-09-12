package datastore

import (
	"context"

	"github.com/jackc/pgx/v5"
)


type DataStore struct {

	conn *pgx.Conn //database connection
	
}


func NewDatastore(ctx context.Context) (*DataStore, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/demo_dev")
	if err != nil {
		return nil, err
	}


	return &DataStore { 
		conn: conn,
	}, nil
	
}

func (s *DataStore) CloseConnection(ctx context.Context) error {
	return s.conn.Close(ctx)
}


func (s *DataStore) SelectEmailByUserID(userID string) (string, error) {
	var email string
	err := s.conn.QueryRow(context.Background(), "SELECT email FROM users WHERE id=$1", userID).Scan(&email)
	if err != nil {
		return email, err
	}

	return email, nil
}
