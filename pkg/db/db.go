package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (database *Database) Connect() error {
	db, err := sql.Open("postgres", "postgres://anton:s3cr3t@127.0.0.1/kertificate?sslmode=disable")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	database.db = db
	return nil
}
