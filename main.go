package main

import (
	"encoding/json"
	"fetchrewards-go/models"
	"fetchrewards-go/services"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := services.ProcessReceipt(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the UUID part from the URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest)
		return
	}
	idStr := pathParts[2] // Assuming URL is "/receipts/{uuid}/points"

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest)
		return
	}

	points, err := services.GetReceiptPoints(id)
	if err != nil {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func setupRoutes() {
	http.HandleFunc("/receipts/process", processReceiptHandler)
	http.HandleFunc("/receipts/", getReceiptPointsHandler) // Adjust path as necessary
}

func main() {
	setupRoutes()
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
