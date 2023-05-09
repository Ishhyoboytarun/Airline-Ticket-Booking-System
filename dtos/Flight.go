package dtos

type AddFlightRequest struct {
	Name  string
	Route *AddRouteRequest
	Configuration
}

type Configuration struct {
	SeatingCapacity int
	Arrangement     []*Seat
}

type SearchFlightRequest struct {
	Source      string
	Destination string
}
