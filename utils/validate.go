package utils

import "encoding/json"

// IsJSON checks whether data is a valid JSON encoding
func IsJSON(data []byte) bool {
	return json.Valid(data)
}

func IsJsonWithInterface(data interface{}) bool {
	var t interface{}
	switch val := data.(type) {
	case string:
		json.Unmarshal([]byte(val), &t)
	default:
		t = data
	}
	_, ok := t.(map[string]interface{})
	return ok
}
