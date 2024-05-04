package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadJSON reads the body of the given request as JSON encoded data and
// unmarshalls it into the 'into' pointer.
func ReadJSON(req *http.Request, dst interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(dst); err != nil {
		return fmt.Errorf("failed to decode body: %s", err.Error())
	}
	return nil
}

// ReadJSON is to return the encoded http.Response struct from server to the client
func WriteJSON(w http.ResponseWriter, r *http.Request, status int, response any) error {
	req := r.WithContext(context.WithValue(r.Context(), StatusCode, status))
	*r = *req
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}
