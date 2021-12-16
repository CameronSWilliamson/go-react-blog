package utils

import (
	"encoding/json"
	"net/http"
)

func GetEncoder(w *http.ResponseWriter) *json.Encoder {
	(*w).Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(*w)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false)
	return encoder
}
