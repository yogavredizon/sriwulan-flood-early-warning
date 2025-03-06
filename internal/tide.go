package internal

import (
	"database/sql"
	"errors"

	"github.com/yogavredizon/sriwulan-flood-early-warning/scrapper"
)

// internal package contains all function to work with database

type Tide struct {
	DB *sql.DB
}

func NewDBTide(s *sql.DB) Tide {
	return Tide{
		DB: s,
	}
}

func (w *Weather) AddTide(tide scrapper.Tide) error {
	query := "INSERT INTO weather VALUES(?, ?, ?)"
	stmt, err := w.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(tide.Time, tide.Data)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
