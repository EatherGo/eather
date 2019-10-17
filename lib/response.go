package lib

// EatherResponse struct - customize response for routes
type EatherResponse struct {
	Type string
	Data interface{}
}

// NewResponse - create new response
func NewResponse(data interface{}) EatherResponse {
	response := EatherResponse{Type: "JSON", Data: data}

	return response
}
