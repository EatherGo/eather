package lib

import (
	"encoding/json"
	"net/http"
)

// EatherResponse struct - customize response for routes
type EatherResponse struct {
	Status  bool
	Message string
	Data    Response
}

// Response set response type
type Response map[string]interface{}

// SendJSONResponse will set type to application/json and send to response
func SendJSONResponse(w http.ResponseWriter, r EatherResponse) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(r)
}
