package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type ContextHeaderKey string

const (
	RequestIDKey    ContextHeaderKey = "RequestID"
	RequestIDHeader string           = "Request-ID"
)

func RequestID() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(RequestIDHeader)

			if requestID == "" {
				fmt.Println("TODO: empty requestId")
				// if isRequired {
				// 	w.Header().Set("Content-Type", "application/json")
				// 	w.WriteHeader(http.StatusBadRequest)
				// 	//TODO replace with proper error code based on standard
				// 	errorResponse := BuildErrorResponse("0001", "Correlation-ID header is required")
				// 	errorResponseBytes, _ := json.Marshal(errorResponse)
				// 	w.Write(errorResponseBytes)
				// 	return
				// }

				//use uuid4 for now
				requestID = uuid.New().String()
			}

			ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
