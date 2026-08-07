package midleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMidleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		log.Printf("[%s] %s %s - %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start))

		next.ServeHTTP(w, r)
	})
}
