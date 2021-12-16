package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cameronswilliamson/go-react-blog/models"
)

// Endpoint that returns all categories
func FetchCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	allCategories, err := models.FetchAllCategories()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	encoder.Encode(allCategories)
}
