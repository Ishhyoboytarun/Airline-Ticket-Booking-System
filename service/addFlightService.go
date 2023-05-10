package service

import (
	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/repo"
)

type AddFlightService struct {
	AddFlightRepo *repo.AddFlightRepo
}

func NewAddFlightService(RegisterUserRepo *repo.AddFlightRepo) *AddFlightService {
	return &AddFlightService{
		AddFlightRepo: RegisterUserRepo,
	}
}

func (r *AddFlightService) AddFlight(flightDetails *dtos.AddFlightRequest) error {

	err := r.AddFlightRepo.AddFlight(flightDetails)
	if err != nil {
		return err
	}
	return nil
}
