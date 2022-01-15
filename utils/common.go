package utils

import "encoding/json"

func Stringify(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
