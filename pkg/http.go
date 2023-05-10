package catalog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartHttpServer starts a new HTTP server
func StartHttpServer() {
	// mux creation
	router := mux.NewRouter()

	// routes definition
	router.HandleFunc("/api/flavours", getAllFlavoursHandler).Methods("GET")
	router.HandleFunc("/api/flavours/architecture/{architecture}", getFlavoursByArchitectureHandler).Methods("GET")
	router.HandleFunc("/api/flavours/os/{os}", getFlavoursByOSHandler).Methods("GET")

	// Avvio del server HTTP
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAllFlavoursHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve Flavours from MongoDB based on the architecture
	flavours, err := getAllFlavours()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavours)
}

func getFlavoursByArchitectureHandler(w http.ResponseWriter, r *http.Request) {
	// Get the architecture value from the URL parameters
	params := mux.Vars(r)
	architecture := params["architecture"]

	// Retrieve Flavours from MongoDB based on the architecture
	flavours, err := getFlavoursByArchitecture(architecture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavours)
}

func getFlavoursByOSHandler(w http.ResponseWriter, r *http.Request) {
	// Get the architecture value from the URL parameters
	params := mux.Vars(r)
	os := params["os"]

	// Retrieve Flavours from MongoDB based on the architecture
	flavours, err := getFlavoursByOS(os)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavours)
}
