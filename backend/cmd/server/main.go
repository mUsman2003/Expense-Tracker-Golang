package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Name string `json:"name"`
}

func getNameHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Name: "Usman"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/name", getNameHandler)

	println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
