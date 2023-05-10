package service

import (
	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/repo"
)

type AddRouteService struct {
	AddRouteRepo *repo.AddRouteRepo
}

func NewAddRouteService(AddRouteRepo *repo.AddRouteRepo) *AddRouteService {
	return &AddRouteService{
		AddRouteRepo: AddRouteRepo,
	}
}

func (r *AddRouteService) AddRoute(userDetails *dtos.AddRouteRequest) error {

	err := r.AddRouteRepo.AddRoute(userDetails)
	if err != nil {
		return err
	}
	return nil
}
