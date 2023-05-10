package repo

import (
	"database/sql"
	"fmt"

	"github.com/dunzoit/Airline-Ticket-Booking-System/dtos"
)

const (
	RegisterUserQuery = `INSERT INTO users (name, email, phone) VALUES ($1, $2, $3)`
)

type RegisterUserRepo struct {
	db *sql.DB
}

func NewRegisterUserRepo(db *sql.DB) *RegisterUserRepo {
	return &RegisterUserRepo{
		db: db,
	}
}

func (r *RegisterUserRepo) RegisterUser(userDetails *dtos.RegisterUserRequest) error {
	result, err := r.db.Exec(RegisterUserQuery, userDetails.Name, userDetails.Email, userDetails.Phone)
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
