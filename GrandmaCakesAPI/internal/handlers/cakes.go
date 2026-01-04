package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var cakes []Cake

type Cake struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Price       float64  `json:"price"`
}

func ListCakes(w http.ResponseWriter, r *http.Request) {
	if len(cakes) == 0 {
		http.Error(w, "No Cakes Found", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cakes)
}
func GetCake(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	for _, cake := range cakes {
		if cake.ID == id {
			json.NewEncoder(w).Encode(cake)
			return
		}
	}

	json.NewEncoder(w).Encode(&Cake{})

}
