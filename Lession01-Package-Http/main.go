package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("Server is starting on port 8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")

		if r.Method != http.MethodGet {
			log.Printf("Unsupported method: %s", r.Method)
			http.Error(w, "This method is not supported", http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/demo", demoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) // Simplified version
}

// demoHandler is a named handler function for "/demo"
// demoHandler returns a JSON response
func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Demo handler hit")

	// Example response data
	response := map[string]string{
		"message": "This is the demo route!",
		"status":  "success",
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-course", "Golang programing")

	// Set status code (optional)
	w.WriteHeader(http.StatusOK)

	// Encode to JSON and send
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
