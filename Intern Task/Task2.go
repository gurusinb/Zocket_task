package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/items", createItem).Methods("POST")


	r.HandleFunc("/items", getItems).Methods("GET")


	r.HandleFunc("/items/{id}", getItem).Methods("GET")


	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")


	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedItem Item
	json.NewDecoder(r.Body).Decode(&updatedItem)
	for i, item := range items {
		if item.ID == params["id"] {
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	json.NewEncoder(w).Encode(http.StatusBadRequest)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range items {
		if item.ID == params["id"] {
			items = append(items[:i], items[i+1:]...)
			json.NewEncoder(w).Encode(http.StatusOK)
			return
		}
	}
	json.NewEncoder(w).Encode(http.StatusBadRequest)