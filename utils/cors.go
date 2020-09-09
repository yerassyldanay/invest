package utils

import (
	http "net/http"
)

func SetCorsHeaders(w *http.ResponseWriter) {
	/*
		Explicitly informs the referer how many seconds it should store the preflight
		result. Within this time, it can just send the request,
		and doesn't need to bother sending the preflight request again.
	 */
	(*w).Header().Set("Access-Control-Max-Age", "86400")

	(*w).Header().Set("Access-Control-Allow-Credentials", "")

	(*w).Header().Set("Access-Control-Allow-Origin", "http://178.170.221.116, http://127.0.0.1:7000")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, EAT")
	(*w).Header().Add("Content-Type", "application/json")
}
