package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/exceptions"
	"github.com/dunzoit/Airline-Ticket-Booking-System/service"
)

type RegisterUser struct {
	RegisterUserService *service.RegisterUserService
}

func NewRegisterUserHandler(RegisterUserService *service.RegisterUserService) *RegisterUser {
	return &RegisterUser{
		RegisterUserService: RegisterUserService,
	}
}

func (r *RegisterUser) RegisterUserHandler(w http.ResponseWriter, request *http.Request) {

	registerUserRequest := &dtos.RegisterUserRequest{}
	if err := json.NewDecoder(request.Body).Decode(&registerUserRequest); err != nil {
		http.Error(w, exceptions.BadRequestError, http.StatusBadRequest)
		return
	}

	err := r.RegisterUserService.RegisterUser(registerUserRequest)
	if err != nil {
		http.Error(w, exceptions.InternalServerError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
