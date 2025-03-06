package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Conn(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
