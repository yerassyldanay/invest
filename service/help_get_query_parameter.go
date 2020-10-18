package service

import (
	"invest/model"
	"net/http"
	"strconv"
)

func OnlyPrepareStatusAndStep(r *http.Request) ([]string, []int) {
	// get query parameters
	var statuses = []string{}
	status := OnlyGetQueryParameter(r, "status", "").(string)
	if status == "" {
		statuses = model.Prepare_project_statuses("")
	} else {
		statuses = []string{status}
	}

	// prepare step
	var steps = []int{}
	step := OnlyGetQueryParameter(r, "step", 0).(int)
	if step == 0 {
		steps = []int{1, 2}
	} else {
		steps = []int{step}
	}

	return statuses, steps
}

func Get_query_parameter_str(r *http.Request, key string, base string) (string) {
	t := r.URL.Query()[key]
	if len(t) == 0 {
		return base
	}

	return t[0]
}

func Get_query_parameter_uint64(r *http.Request, key string, default_value uint64) (uint64) {
	t := r.URL.Query()[key]
	if len(t) != 0 {
		i, err := strconv.ParseInt(t[0], 0, 16)
		if err == nil {
			return uint64(i)
		}
	}

	return default_value
}

func OnlyGetQueryParameterAsList(r *http.Request, key string) ([]string) {
	return r.URL.Query()[key]
}

func OnlyGetQueryParameter(r *http.Request, key string, defval interface{}) interface{} {
	values := r.URL.Query()[key]
	if len(values) < 1 {
		return defval
	}

	t := values[0]

	if t == "" {
		return defval
	}

	/*
		string & bool values are parsed here
	*/
	switch defval.(type) {
	case bool:
		i, err := strconv.ParseBool(t)
		if err != nil {
			return defval
		}
		return i
	case string:
		return t
	}

	/*
		will parse as integer then encapsulate the value
	*/
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return defval
	}

	switch defval.(type) {
	case int:
		return int(i)
	case int8:
		return int8(i)
	case int16:
		return int16(i)
	case int32:
		return int32(i)
	case int64:
		return int64(i)
	case uint:
		return uint(i)
	case uint8:
		return uint8(i)
	case uint16:
		return uint16(i)
	case uint32:
		return uint32(i)
	case uint64:
		return uint64(i)
	}

	return defval
}

/*
	this function will get header value for you
		the type of the value that will be returned depends on the type of a value provided
			e.g. if defval is uint64 then uint64 value will be returned
 */
func Get_header_parameter(r *http.Request, key string, defval interface{}) interface{} {
	t := r.Header.Get(key)

	if t == "" {
		return defval
	}

	/*
		string & bool values are parsed here
	 */
	switch defval.(type) {
	case bool:
		i, err := strconv.ParseBool(t)
		if err != nil {
			return defval
		}
		return i
	case string:
		return t
	}

	/*
		will parse as integer then encapsulate the value
	*/
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return defval
	}

	switch defval.(type) {
	case int:
		return int(i)
	case int8:
		return int8(i)
	case int16:
		return int16(i)
	case int32:
		return int32(i)
	case int64:
		return int64(i)
	case uint:
		return uint(i)
	case uint8:
		return uint8(i)
	case uint16:
		return uint16(i)
	case uint32:
		return uint32(i)
	case uint64:
		return uint64(i)
	}

	return defval
}

func OnlyConvertString(val string, defval interface{}) interface{} {
	if val != "" {
		return OnlyConvert([]string{val}, defval)
	}
	return defval
}

func OnlyConvert(values []string, defval interface{}) interface{} {
	if len(values) < 1 {
		return defval
	}

	t := values[0]

	if t == "" {
		return defval
	}

	/*
		string & bool values are parsed here
	*/
	switch defval.(type) {
	case bool:
		i, err := strconv.ParseBool(t)
		if err != nil {
			return defval
		}
		return i
	case string:
		return t
	}

	/*
		will parse as integer then encapsulate the value
	*/
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return defval
	}

	switch defval.(type) {
	case int:
		return int(i)
	case int8:
		return int8(i)
	case int16:
		return int16(i)
	case int32:
		return int32(i)
	case int64:
		return int64(i)
	case uint:
		return uint(i)
	case uint8:
		return uint8(i)
	case uint16:
		return uint16(i)
	case uint32:
		return uint32(i)
	case uint64:
		return uint64(i)
	}

	return defval
}