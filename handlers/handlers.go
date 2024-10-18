package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var urlMap = make(map[string]string)

// SearchHandler HomeHandler handles the root endpoint
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tinyURL := strings.TrimPrefix(r.URL.Path, "/")
	longURL, exists := urlMap[tinyURL]
	if !exists {
		http.Error(w, "TinyURL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}

// ResourceHandler handles CRUD operations for a resource
func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createResource(w, r)
	case http.MethodPut:
		updateResource(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// createResource handles POST requests
func createResource(w http.ResponseWriter, r *http.Request) {
	// Logic to create resource
	var newResource map[string]string
	err := json.NewDecoder(r.Body).Decode(&newResource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tinyURL := generateTinyUrl(newResource)
	URL, exists := newResource["url"]
	if !exists || URL == "" {
		http.Error(w, "Missing field: url", http.StatusBadRequest)
		return
	}
	response := map[string]string{
		"key":       tinyURL,
		"long_url":  newResource["url"],
		"short_url": fmt.Sprintf("https://lnkd.in/%s", tinyURL),
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generateTinyUrl(resource map[string]string) string {
	longURL := resource["url"]
	hash := sha256.Sum256([]byte(longURL))
	hashStr := hex.EncodeToString(hash[:])
	tinyURL := hashStr[:6]
	urlMap[tinyURL] = longURL
	return tinyURL
}

// updateResource handles PUT requests
func updateResource(w http.ResponseWriter, r *http.Request) {
	// Logic to update resource
	var updatedResource map[string]string
	err := json.NewDecoder(r.Body).Decode(&updatedResource)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedResource)
}
