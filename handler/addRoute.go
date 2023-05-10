package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/exceptions"
	"github.com/dunzoit/Airline-Ticket-Booking-System/service"
)

type AddRoute struct {
	AddRouteService *service.AddRouteService
}

func NewAddRouteHandler(AddRouteService *service.AddRouteService) *AddRoute {
	return &AddRoute{
		AddRouteService: AddRouteService,
	}
}

func (r *AddRoute) AddRouteHandler(w http.ResponseWriter, request *http.Request) {

	addRouteRequest := &dtos.AddRouteRequest{}
	if err := json.NewDecoder(request.Body).Decode(&addRouteRequest); err != nil {
		http.Error(w, exceptions.BadRequestError, http.StatusBadRequest)
		return
	}

	err := r.AddRouteService.AddRoute(addRouteRequest)
	if err != nil {
		http.Error(w, exceptions.InternalServerError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
