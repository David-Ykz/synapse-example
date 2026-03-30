package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Flight struct {
	Airline string  `json:"airline"`
	Price   float64 `json:"price"`
}

func HandleGetFlights(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")

	flights := []Flight{
		{Airline: "Air Canada", Price: 450.00},
		{Airline: "United Airlines", Price: 425.50},
		{Airline: "Porter Airlines", Price: 395.99},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flights); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
