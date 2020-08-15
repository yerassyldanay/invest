package utils

import (
	"context"
	"net/http"
	"strconv"
)

// Get retrieves a value from the request context
func GetContext(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

// Set stores a value on the request context
func SetContext(r *http.Request, key, val interface{}) *http.Request {
	if val != nil {
		r.WithContext(context.WithValue(r.Context(), key, val))
	}

	return r
}

/*
	get & set header
 */
func SetHeader(r *http.Request, key, val string) *http.Request {
	r.Header.Set(key, val)
	return r
}

func GetHeader(r *http.Request, key string) uint64 {
	if t := r.Header.Get(key); t != "" {
		a, _ := strconv.ParseInt(t, 0, 10)
		return uint64(a)
	}
	return 0
}