package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//go:generate counterfeiter . Database
type Database interface {
	Ping() error
	Query(string, ...interface{}) (Rows, error)
}

type PostgresDatabase struct {
	db *sql.DB
}

func New(connInfo string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{
		db: db,
	}, nil
}

func (p *PostgresDatabase) Ping() error {
	return p.db.Ping()
}

func (p *PostgresDatabase) Query(query string, args ...interface{}) (Rows, error) {
	return p.db.Query(query, args...)
}

//go:generate counterfeiter . Rows
type Rows interface {
	Next() bool
	Scan(...interface{}) error
}
