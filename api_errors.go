package apierrors

import (
	"encoding/json"
	"net/http"

	"github.com/joeyfromspace/go-api-errors/v2/errors"
)

// APIErrorResponse is the envelope for a json error response
type APIErrorResponse struct {
	Error *errors.APIError `json:"error"`
}

// SendError sends an error response to the passed in writer along with an optional custom error struct
func SendError(w http.ResponseWriter, err *errors.APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)

	j := json.NewEncoder(w)
	j.Encode(err)
}

// SendInternalError sends a 500 response to the passed in writer along with an optional custom error struct
func SendInternalError(w http.ResponseWriter, e *errors.APIError) {
	err := errors.NewInternalError(e)
	SendError(w, err)
}
