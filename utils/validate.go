package utils

import "encoding/json"

// IsJSON checks whether data is a valid JSON encoding
func IsJSON(data []byte) bool {
	return json.Valid(data)
}
