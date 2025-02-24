package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Item struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)
    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}

func main() {
    router := mux.NewRouter()

    items = append(items, Item{ID: "1", Name: "Sample Item", Description: "This is a sample item"})

    router.HandleFunc("/api/items", getItems).Methods("GET")
    router.HandleFunc("/api/items", createItem).Methods("POST")

    http.ListenAndServe(":8080", router)
}
