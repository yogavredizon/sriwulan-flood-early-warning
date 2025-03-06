package internal

import (
	"database/sql"
	"errors"

	"github.com/yogavredizon/sriwulan-flood-early-warning/scrapper"
)

// internal package contains all function to work with database

type Weather struct {
	DB *sql.DB
}

func NewDBWeather(s *sql.DB) Weather {
	return Weather{
		DB: s,
	}
}

func (w *Weather) AddWeather(weather scrapper.Weather) error {
	query := "INSERT INTO weather VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := w.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(weather.Time, weather.Image, weather.Temperature, weather.State, weather.Humidity, weather.WindSpeed, weather.WindDirection, weather.Visibility, weather.LastUpdate)

	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
