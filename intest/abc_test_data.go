package intest

import (
	"bytes"
	"fmt"
	"net/url"
)

type TestData struct {
	URL				string
	Query			url.Values

	Body			map[string]interface{}
	Method			string

	Resp			map[string]interface{}
	RespStatus		int
}

func (t *TestData) Convert_body_to_string() string {
	var b bytes.Buffer
	var l = len(t.Body)

	for key, value := range t.Body {
		b.WriteString(fmt.Sprintf(`"%s":"%v"`, key, value))

		if l > 1 { b.WriteString(", ") }

		l--
	}

	var result = fmt.Sprintf(`{%s}`, b.String())
	//fmt.Println("Convert_body_to_string: ", result)
	return result
}

func (t *TestData) Add_query_parameters() {
	if q := t.Query.Encode(); q != "" {
		t.URL = t.URL + "?" + t.Query.Encode()
	}
}
