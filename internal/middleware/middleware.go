package middleware

import (
	"context"
	"net/http"
	"uuid"
)

type contextKey string

const XRequestIDKey contextKey = "X-Request-ID"

func SetRequestUUIDToHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.NewV4()
		w.Header().Set(string(XRequestIDKey), uuid.String())
		ctx := context.WithValue(r.Context(), XRequestIDKey, uuid.String())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(XRequestIDKey).(string)
	if !ok {
		return ""
	}

	return requestID
}
