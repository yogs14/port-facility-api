package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
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
	
	// Perbaikan G104: Cek error saat mengubah data ke JSON
	err := json.NewEncoder(w).Encode(GetCraneStatus())
	if err != nil {
		// Jika gagal, kirim status 500 (Internal Server Error)
		http.Error(w, "Gagal memproses data", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Membuat router HTTP
	mux := http.NewServeMux()
	mux.HandleFunc("/api/status", statusHandler)

	// Perbaikan G114: Membuat server kustom dengan batas waktu (Timeout)
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,  // Maksimal waktu baca header
		ReadTimeout:       5 * time.Second,  // Maksimal waktu baca seluruh request
		WriteTimeout:      10 * time.Second, // Maksimal waktu untuk membalas request
		IdleTimeout:       120 * time.Second,// Waktu tunggu maksimal untuk koneksi yang nganggur
	}

	log.Println("Server berjalan di port 8080...")
	
	// Perbaikan G104: Menangkap error jika server gagal menyala
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server berhenti mendadak: %v", err)
	}
}