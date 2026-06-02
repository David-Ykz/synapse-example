package main

import (
	"log"
	"net/http"

	handlers "synapse-example/handlers/internal"
)

func main() {
	http.HandleFunc("/get_flights", handlers.HandleGetFlights)
	http.HandleFunc("/get_hotels", handlers.HandleGetHotels)
	http.HandleFunc("/get_restaurants", handlers.HandleGetRestaurants)

	log.Println("Server starting on :8000...")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
