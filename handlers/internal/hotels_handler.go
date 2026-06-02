package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hotel struct {
	Name         string  `json:"name"`
	PricePerNight float64 `json:"price_per_night"`
}

var hotelsByCity = map[string][]Hotel{
	"zurich": {
		{Name: "The Dolder Grand", PricePerNight: 680.00},
		{Name: "Hotel Baur au Lac", PricePerNight: 720.00},
		{Name: "25hours Hotel Zürich West", PricePerNight: 210.00},
	},
	"geneva": {
		{Name: "Four Seasons Hotel des Bergues", PricePerNight: 890.00},
		{Name: "Hôtel Beau-Rivage", PricePerNight: 540.00},
		{Name: "Hotel N'vY", PricePerNight: 195.00},
	},
	"bern": {
		{Name: "Bellevue Palace", PricePerNight: 430.00},
		{Name: "Hotel Schweizerhof Bern", PricePerNight: 310.00},
		{Name: "Hotel Allegro", PricePerNight: 245.00},
	},
}

func HandleGetHotels(w http.ResponseWriter, r *http.Request) {
	log.Println("Received get_hotels request")

	city := r.URL.Query().Get("city")

	hotels, ok := hotelsByCity[city]
	if !ok {
		http.Error(w, "No hotels found for city: "+city, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(hotels); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
