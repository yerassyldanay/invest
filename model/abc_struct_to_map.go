package model

import (
	"reflect"
	"strings"
)

/*
	convert struct to map format that
		it can be sent as a map
 */
func Struct_to_map(cs interface{}) (map[string]interface{}) {
	var info = make(map[string]interface{})
	v := reflect.ValueOf(cs)
	typeOfS := v.Type()

	if v.NumField() < 1 {
		return nil
	}

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(string(typeOfS.Field(i).Tag.Get("json")))
		if key == "" {
			key = strings.ToLower(typeOfS.Field(i).Name)
		}
		info[key] = v.Field(i).Interface()
	}
	//b, err := json.Marshal(cs)
	//if err != nil {
	//	return info
	//}
	//
	//_ = json.Unmarshal(b, &info)
	return info
}

func Struct_to_map_with_escape(cs interface{}, escape []string) (map[string]interface{}) {
	var info = make(map[string]interface{})
	v := reflect.ValueOf(cs)
	typeOfS := v.Type()

	if v.NumField() < 1 {
		return nil
	}

	var ok = true
	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(typeOfS.Field(i).Tag.Get("json"))

		if key == "-" {
			continue
		}

		ok = true
		for _, tkey := range escape {
			if key == tkey {
				ok = false
				break
			}
		}

		if !ok {
			continue
		}

		info[key] = v.Field(i).Interface()
	}

	return info
}

func List_of_struct_to_map(arr []interface{}) ([]map[string]interface{}) {
	var result = []map[string]interface{}{}
	for _, elem := range arr {
		result = append(result, Struct_to_map(elem))
	}

	return result
}
