package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type healthCheckResponse struct {
	Test string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	healthCheckResult := healthCheckResponse{
		Test: http.StatusText(http.StatusOK),
	}

	json.NewEncoder(w).Encode(healthCheckResult)
}

func main() {
	// Connecting route to health check handler
	http.HandleFunc("/health", healthCheckHandler)

	// Starting the server on port 8080
	fmt.Println("Starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
