package handlers

import (
	"encoding/json"
	"net/http"
)

type Schedule struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListSchedules(w http.ResponseWriter, _ *http.Request) {
	items := []Schedule{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
