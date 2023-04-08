package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// TODO: chi logger zerolog implementation

func MiddlewareLogger(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	mw := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger,
		NoColor: false,
	})
	return mw
}

func MiddlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(middleware.RequestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		ctx = context.WithValue(ctx, middleware.RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	return middleware.GetReqID(ctx)
}
