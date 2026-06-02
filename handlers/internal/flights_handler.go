package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Flight struct {
	Airline string  `json:"airline"`
	Price   float64 `json:"price"`
}

type priceRange struct{ min, max float64 }

// Mock price ranges for each airport
var rangeByArrival = map[string]priceRange{
	"ZRH": {900, 1400}, // Zurich
	"GVA": {780, 1200}, // Geneva
	"BRN": {620, 950},  // Bern
}

var defaultRange = priceRange{700, 1100}

// Scale price by month
var seasonalFactors = [12]float64{
	0.85, 0.80, 0.90, 0.95, 1.00, 1.10,
	1.30, 1.25, 1.05, 0.95, 0.85, 1.20,
}

func randomInRange(r priceRange) float64 {
	raw := r.min + rand.Float64()*(r.max-r.min)
	return float64(int(raw*100)) / 100
}

func HandleGetFlights(w http.ResponseWriter, r *http.Request) {
	arrival := strings.ToUpper(r.URL.Query().Get("arrival"))
	log.Printf("Received get_flights request (arrival=%s)", arrival)

	pr, ok := rangeByArrival[arrival]
	if !ok {
		pr = defaultRange
	}
	seasonal := seasonalFactors[time.Now().Month()-1]

	scaled := priceRange{
		min: pr.min * seasonal,
		max: pr.max * seasonal,
	}

	flights := []Flight{
		{Airline: "Swiss International Air Lines", Price: randomInRange(scaled)},
		{Airline: "Air Canada", Price: randomInRange(scaled)},
		{Airline: "Lufthansa", Price: randomInRange(scaled)},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flights); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
