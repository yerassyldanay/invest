package helper

import (
	"encoding/json"
	"fmt"
)

func HelperPrint(any interface{}) {
	b, err := json.MarshalIndent(any, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
