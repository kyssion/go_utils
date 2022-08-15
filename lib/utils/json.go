package util

import "encoding/json"

func CheckJSON(str string, data interface{}) error {
	return json.Unmarshal([]byte(str), &data)
}
