package service

import (
	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
	"github.com/dunzoit/Airline-Ticket-Booking-System/repo"
)

type RegisterUserService struct {
	RegisterUserRepo *repo.RegisterUserRepo
}

func NewRegisterUserService(RegisterUserRepo *repo.RegisterUserRepo) *RegisterUserService {
	return &RegisterUserService{
		RegisterUserRepo: RegisterUserRepo,
	}
}

func (r *RegisterUserService) RegisterUser(userDetails *dtos.RegisterUserRequest) error {

	err := r.RegisterUserRepo.RegisterUser(userDetails)
	if err != nil {
		return err
	}
	return nil
}
