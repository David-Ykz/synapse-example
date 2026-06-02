package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Restaurant struct {
	Name         string  `json:"name"`
	Cuisine      string  `json:"cuisine"`
	Rating       float64 `json:"rating"`
	AveragePrice float64 `json:"average_price"`
}

var restaurantsByCity = map[string][]Restaurant{
	"zurich": {
		{Name: "Restaurant Kronenhalle", Cuisine: "Swiss", Rating: 4.6, AveragePrice: 95.00},
		{Name: "Hiltl", Cuisine: "Vegetarian", Rating: 4.5, AveragePrice: 40.00},
		{Name: "Pavillon", Cuisine: "French", Rating: 4.8, AveragePrice: 180.00},
	},
	"geneva": {
		{Name: "Le Chat-Botté", Cuisine: "French", Rating: 4.7, AveragePrice: 160.00},
		{Name: "Brasserie Lipp", Cuisine: "Swiss-French", Rating: 4.4, AveragePrice: 65.00},
		{Name: "Rasoi by Vineet", Cuisine: "Indian", Rating: 4.6, AveragePrice: 120.00},
	},
	"bern": {
		{Name: "Restaurant Zimmermania", Cuisine: "Swiss", Rating: 4.5, AveragePrice: 55.00},
		{Name: "Lorenzini", Cuisine: "Italian", Rating: 4.4, AveragePrice: 50.00},
		{Name: "Jack's Brasserie", Cuisine: "French", Rating: 4.6, AveragePrice: 75.00},
	},
}

func HandleGetRestaurants(w http.ResponseWriter, r *http.Request) {
	city := strings.ToLower(r.URL.Query().Get("city"))
	log.Printf("Received get_restaurants request (city=%s)", city)

	restaurants, ok := restaurantsByCity[city]
	if !ok {
		http.Error(w, "No restaurants found for city: "+city, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(restaurants); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
