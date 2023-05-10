package models

type Booking struct {
	Id          int
	Customers   []*Customer
	FlightId    int
	Seats       []*Seat
	FlightName  string
	SeatCount   int
	Source      string
	Destination string
	Amount      int
}
