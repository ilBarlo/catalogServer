package catalog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// StartHttpServer starts a new HTTP server
func StartHttpServer() {
	// mux creation
	router := mux.NewRouter()

	// routes definition
	router.HandleFunc("/api/flavours", getAllFlavoursHandler).Methods("GET")
	router.HandleFunc("/api/flavours/request", findMatchingResourcesHandler).Methods("GET")
	router.HandleFunc("/api/flavours/architecture/{architecture}", getFlavoursByArchitectureHandler).Methods("GET")
	router.HandleFunc("/api/flavours/os/{os}", getFlavoursByOSHandler).Methods("GET")

	// Start server HTTP
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// getAllFlavoursHandler handles the request to get all Flavours
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

// getFlavoursByArchitectureHandler handles the request to get all Flavours matching an Archichecture
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

// getFlavoursByOSHandler handles the request to get all Flavours matching an OS
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

// findMatchingResourcesHandler handles the request to get all Flavours matching a specific request
func findMatchingResourcesHandler(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	// Estrarre i valori dei parametri desiderati dalla mappa "params"
	request := Request{
		Architecture:  params.Get("architecture"),
		OS:            params.Get("os"),
		CPURequest:    params.Get("cpuRequest"),
		MemoryRequest: params.Get("memoryRequest"),
		PodsPlan:      strings.ReplaceAll(params.Get("podsPlan"), "\n", ""),
	}

	// Retrieve Flavours from MongoDB based on the architecture
	flavours, err := findMatchingResources(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the Flavours as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flavours)

}
