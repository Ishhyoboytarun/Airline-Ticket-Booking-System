package main

import (
	"database/sql"
	"net/http"

	"github.com/dunzoit/Airline-Ticket-Booking-System/handler"
	"github.com/dunzoit/Airline-Ticket-Booking-System/repo"
	"github.com/dunzoit/Airline-Ticket-Booking-System/service"
	"github.com/gorilla/mux"
)

type Dependencies struct {
	RegisterUserHandler *handler.RegisterUser
	SearchFlightHandler *handler.SearchFlight
	AddRouteHandler     *handler.AddRoute
	AddFlightHandler    *handler.AddFlight
}

func main() {
	router := mux.NewRouter()
	handler := initializeDependencies()

	router.HandleFunc("/register", handler.RegisterUserHandler.RegisterUserHandler).Methods(http.MethodPost)

	router.HandleFunc("/add-flight", handler.AddFlightHandler.AddFlightHandler).Methods(http.MethodPost)

	router.HandleFunc("/add-route", handler.AddRouteHandler.AddRouteHandler).Methods(http.MethodPost)

	router.HandleFunc("/search-flight", handler.SearchFlightHandler.SearchFlightHandler).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}

func initializeDependencies() *Dependencies {
	db := &sql.DB{}
	registerUserRepo := repo.NewRegisterUserRepo(db)
	searchFlightRepo := repo.NewSearchFlightRepo(db)
	addRouteRepo := repo.NewAddRouteRepo(db)
	addFlightRepo := repo.NewAddFlightRepo(db)
	registerUserService := service.NewRegisterUserService(registerUserRepo)
	searchFlightService := service.NewSearchFlightService(searchFlightRepo)
	addRouteService := service.NewAddRouteService(addRouteRepo)
	addFlightService := service.NewAddFlightService(addFlightRepo)
	registerUserHandler := handler.NewRegisterUserHandler(registerUserService)
	searchFlightHandler := handler.NewSearchFlightHandler(searchFlightService)
	addRouteHandler := handler.NewAddRouteHandler(addRouteService)
	addFlightHandler := handler.NewAddFlightHandler(addFlightService)
	return &Dependencies{
		RegisterUserHandler: registerUserHandler,
		SearchFlightHandler: searchFlightHandler,
		AddRouteHandler:     addRouteHandler,
		AddFlightHandler:    addFlightHandler,
	}
}
