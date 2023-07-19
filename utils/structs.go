package utils

import (
	"reflect"
)

func StructToStruct(sourceStruct interface{}, targetStruct interface{}) {
	source := structToMap(sourceStruct)
	targetV := reflect.ValueOf(targetStruct)
	targetT := reflect.TypeOf(targetStruct)
	if targetV.Kind() == reflect.Ptr {
		targetV = targetV.Elem()
		targetT = targetT.Elem()
	}
	for i := 0; i < targetV.NumField(); i++ {
		fieldName := targetT.Field(i).Name
		sourceVal := source[fieldName]
		if !sourceVal.IsValid() {
			continue
		}
		targetVal := targetV.Field(i)
		if targetVal.Type() != sourceVal.Type() {
			continue
		}
		targetVal.Set(sourceVal)
	}
}

func structToMap(structName interface{}) map[string]reflect.Value {
	t := reflect.TypeOf(structName)
	v := reflect.ValueOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	fieldNum := t.NumField()
	resMap := make(map[string]reflect.Value, fieldNum)
	for i := 0; i < fieldNum; i++ {
		resMap[t.Field(i).Name] = v.Field(i)
	}
	return resMap
}

func isSliceOrArray(ptr interface{}) bool {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		return false
	}
	elemKind := rv.Elem().Kind()
	return elemKind == reflect.Slice || elemKind == reflect.Array
}
