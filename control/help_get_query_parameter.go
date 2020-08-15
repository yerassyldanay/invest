package control

import (
	"net/http"
	"strconv"
)

func Get_query_parameter_str(r *http.Request, key string, base string) (string) {
	t := r.URL.Query()[key]
	if len(t) == 0 {
		return base
	}

	return t[0]
}

func Get_query_parameter_int(r *http.Request, key string, base int) (int) {
	t := r.URL.Query()[key]
	if len(t) != 0 {
		i, err := strconv.ParseInt(t[0], 0, 16)
		if err == nil {
			return int(i)
		}
	}

	return base
}