package repo

import (
	"database/sql"
	"fmt"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
)

const (
	AddFlightQuery = `INSERT INTO flights (name, source, destination) VALUES ($1, $2, $3)`
)

type AddFlightRepo struct {
	db *sql.DB
}

func NewAddFlightRepo(db *sql.DB) *AddFlightRepo {
	return &AddFlightRepo{
		db: db,
	}
}

func (r *AddFlightRepo) AddFlight(flightDetails *dtos.AddFlightRequest) error {
	result, err := r.db.Exec(AddFlightQuery, flightDetails.Name, flightDetails.Route.Source,
		flightDetails.Route.Destination)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("%v no. of rows are affected", rows)
		return err
	}
	return nil
}
