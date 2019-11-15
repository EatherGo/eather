package eather

import (
	"encoding/json"
	"net/http"
)

// Response struct - customize response for routes
type Response struct {
	Status     bool
	Message    string
	Data       DataResponse
	StatusCode int
}

// DataResponse set DataResponse type
type DataResponse map[string]interface{}

// SendJSONResponse will set type to application/json and send to response
func SendJSONResponse(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")

	if !r.Status {
		if r.StatusCode == 0 {
			r.StatusCode = http.StatusBadRequest
		}

		http.Error(w, r.Message, r.StatusCode)
		return
	}

	json.NewEncoder(w).Encode(r)
}
