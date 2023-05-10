package repo

import (
	"database/sql"
	"fmt"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
)

const (
	AddRouteQuery = `INSERT INTO routes (source, destination) VALUES ($1, $2)`
)

type AddRouteRepo struct {
	db *sql.DB
}

func NewAddRouteRepo(db *sql.DB) *AddRouteRepo {
	return &AddRouteRepo{
		db: db,
	}
}

func (r *AddRouteRepo) AddRoute(routeDetails *dtos.AddRouteRequest) error {
	result, err := r.db.Exec(AddRouteQuery, routeDetails.Source, routeDetails.Destination)
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
