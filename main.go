package main

import (
	"encoding/json"
	"net/http"
)

type FacilityStatus struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func GetCraneStatus() FacilityStatus {
	return FacilityStatus{
		ID:     "CRN-001",
		Name:   "Container Crane Alpha",
		Status: "OPERATIONAL",
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetCraneStatus())
}

func main() {
	http.HandleFunc("/api/status", statusHandler)
	http.ListenAndServe(":8080", nil)
}