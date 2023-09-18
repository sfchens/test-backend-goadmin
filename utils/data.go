package utils

import (
	"encoding/json"
	"strconv"
)

type DataTmp struct {
	Label string
	Value interface{}
}

// Bytes2Interface parses the given data into an interface
func Bytes2Interface(jsonBlob []byte) interface{} {
	var v interface{}
	if IsJSON(jsonBlob) {
		_ = json.Unmarshal(jsonBlob, &v)
	}

	return v
}

func IntToStringArray(intArr []int) (parentIdStr []string) {
	for _, val := range intArr {
		res := strconv.Itoa(val)
		parentIdStr = append(parentIdStr, res)
	}
	return
}

func StringToIntArray(stringArr []string) (intArr []int) {
	for _, val := range stringArr {
		i, _ := strconv.Atoi(val)
		intArr = append(intArr, i)
	}
	return
}

func ToJson(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func JsonToInterface(data interface{}, res interface{}) {
	t, _ := json.Marshal(data)
	json.Unmarshal(t, &res)
}

func IsTypeDefaultValue(typeT interface{}) bool {

	switch typeT.(type) {
	case int:
		val, _ := typeT.(int)
		return val == 0
	case int32:
		val, _ := typeT.(int)
		return val == 0
	case int64:
		val, _ := typeT.(int)
		return val == 0
	default:
		return false
	}
}
