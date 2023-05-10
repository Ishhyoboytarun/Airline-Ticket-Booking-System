package models

type Flight struct {
	Id             int
	Name           string
	Route          *Route
	Capacity       int
	AvailableSeats []*Seat
}
