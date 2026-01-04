package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var clients []Client

type Client struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ListClients(w http.ResponseWriter, r *http.Request) {
	if len(clients) == 0 {
		http.Error(w, "No Clients Found", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(clients)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for _, client := range clients {
		if client.ID == id {
			json.NewEncoder(w).Encode(client.Name)
			return
		}
	}

	http.Error(w, "client not found", http.StatusNotFound)

}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client Client

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	client.ID = strconv.Itoa(len(clients) + 1)
	clients = append(clients, client)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updated Client
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	for i, client := range clients {
		if client.ID == id {
			updated.ID = id // URL is source of truth
			clients[i] = updated

			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "client not found", http.StatusNotFound)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, client := range clients {
		if client.ID == id {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			w.WriteHeader(http.StatusNoContent)
			break
		}
	}
	http.Error(w, "client not found", http.StatusNotFound)

}
