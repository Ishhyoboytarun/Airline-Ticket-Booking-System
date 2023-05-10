package service

import (
	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/models"
	"github.com/dunzoit/Airline-Ticket-Booking-System/repo"
)

type SearchFlightService struct {
	SearchFlightRepo *repo.SearchFlightRepo
}

func NewSearchFlightService(SearchFlightRepo *repo.SearchFlightRepo) *SearchFlightService {
	return &SearchFlightService{
		SearchFlightRepo: SearchFlightRepo,
	}
}

func (s *SearchFlightService) SearchFlights(searchFlightDetails *dtos.SearchFlightRequest) (*models.RouteToFlightMapping,
	error) {

	response, err := s.SearchFlightRepo.SearchFlights(searchFlightDetails)
	if err != nil {
		return nil, err
	}
	return response, nil
}
