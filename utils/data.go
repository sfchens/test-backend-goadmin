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
