package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - [%s] \"%s %s %s %s\"\n",
			r.Host,
			time.Now().Format(time.RFC1123),
			r.Method,
			r.URL,
			r.Proto,
			r.UserAgent())
		next.ServeHTTP(w, r)
	})
}