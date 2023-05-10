package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/exceptions"
	"github.com/dunzoit/Airline-Ticket-Booking-System/service"
)

type AddFlight struct {
	AddFlightService *service.AddFlightService
}

func NewAddFlightHandler(AddFlightService *service.AddFlightService) *AddFlight {
	return &AddFlight{
		AddFlightService: AddFlightService,
	}
}

func (r *AddFlight) AddFlightHandler(w http.ResponseWriter, request *http.Request) {

	addFlightRequest := &dtos.AddFlightRequest{}
	if err := json.NewDecoder(request.Body).Decode(&addFlightRequest); err != nil {
		http.Error(w, exceptions.BadRequestError, http.StatusBadRequest)
		return
	}

	err := r.AddFlightService.AddFlight(addFlightRequest)
	if err != nil {
		http.Error(w, exceptions.InternalServerError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
