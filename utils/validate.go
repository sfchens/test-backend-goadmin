package utils

import (
	"encoding/json"
	"regexp"
)

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

func IsPhone(phone string) bool {
	pattern := "^1[0-9]{10}$"
	return regexpMatch(pattern, phone)
}

// IsEmail 验证邮箱
func IsEmail(s string) bool {
	pattern := "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	return regexpMatch(pattern, s)
}

func regexpMatch(rule, matchStr string) bool {
	return regexp.MustCompile(rule).MatchString(matchStr)
}
