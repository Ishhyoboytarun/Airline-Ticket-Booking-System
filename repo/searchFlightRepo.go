package repo

import (
	"database/sql"
	"fmt"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/models"
)

const (
	SearchFlightQuery = `SELECT * FROM RouteToFlightMapping where source = $1 and destination = $2`
)

type SearchFlightRepo struct {
	db *sql.DB
}

func NewSearchFlightRepo(db *sql.DB) *SearchFlightRepo {
	return &SearchFlightRepo{
		db: db,
	}
}

func (r *SearchFlightRepo) SearchFlights(searchFlightDetails *dtos.SearchFlightRequest) (*models.RouteToFlightMapping,
	error) {
	rows, err := r.db.Query(SearchFlightQuery, searchFlightDetails.Source, searchFlightDetails.Destination)
	if err != nil {
		return nil, err
	}

	response := &models.RouteToFlightMapping{}
	for rows.Next() {
		var column1Value string
		var column2Value int
		err := rows.Scan(&column1Value, &column2Value)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			return nil, err
		}
		fmt.Println("column1:", column1Value, "column2:", column2Value)
	}
	return response, nil
}
