package middlewares

import (
	"language-learning/internal/api/limiter"
	"net/http"
)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl := limiter.GetRateHandler()
		if rl == nil {
			panic("could not get limiter")
		}

		if rl.UseToken() {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusTooManyRequests)

	})
}
