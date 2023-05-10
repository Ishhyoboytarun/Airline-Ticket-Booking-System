package models

type RouteToFlightMapping struct {
	Id      int
	Route   *Route
	Flights []*Flight
}
