package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/exceptions"
	"github.com/dunzoit/Airline-Ticket-Booking-System/service"
)

type SearchFlight struct {
	SearchFlightService *service.SearchFlightService
}

func NewSearchFlightHandler(SearchFlightService *service.SearchFlightService) *SearchFlight {
	return &SearchFlight{
		SearchFlightService: SearchFlightService,
	}
}

func (s *SearchFlight) SearchFlightHandler(w http.ResponseWriter, request *http.Request) {

	searchFlightRequest := &dtos.SearchFlightRequest{}
	if err := json.NewDecoder(request.Body).Decode(&searchFlightRequest); err != nil {
		http.Error(w, exceptions.BadRequestError, http.StatusBadRequest)
		return
	}

	resp, err := s.SearchFlightService.SearchFlights(searchFlightRequest)
	if err != nil {
		http.Error(w, exceptions.InternalServerError, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, exceptions.MarshalError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
