package api

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func MiddlewareLogger(log *zerolog.Logger) func(next http.Handler) http.Handler {
	// Original Source: https://github.com/ironstar-io/chizerolog
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				t2 := time.Now()

				// Recover and record stack traces in case of a panic
				if rec := recover(); rec != nil {
					log.Error().
						Str("type", "error").
						Timestamp().
						Interface("recover_info", rec).
						Bytes("debug_stack", debug.Stack()).
						Msg("system error")
					http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}

				fields := make(map[string]any, 20)
				if reqID := GetRequestID(r.Context()); reqID != "" {
					fields["request_id"] = reqID
				}
				fields["remote_ip"] = r.RemoteAddr
				fields["url"] = r.URL.Path
				fields["proto"] = r.Proto
				fields["method"] = r.Method
				fields["user_agent"] = r.Header.Get("User-Agent")
				fields["status"] = ww.Status()
				fields["latency_ms"] = float64(t2.Sub(t1).Nanoseconds()) / 1e6
				if bytesIn := r.Header.Get("Content-Length"); bytesIn != "" {
					fields["bytes_in"] = r.Header.Get("Content-Length")
				}
				fields["bytes_out"] = ww.BytesWritten()

				// log end request
				log.Info().
					Str("type", "access").
					Timestamp().
					Fields(fields).
					Msg("incoming request")
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
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
