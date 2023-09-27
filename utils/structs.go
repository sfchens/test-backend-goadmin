package utils

import (
	"reflect"
	"strconv"
)

func StructToStruct(sourceStruct interface{}, targetStruct interface{}) {
	structToStructSetDefault(sourceStruct, targetStruct, false)
}

func StructToStructDefault(sourceStruct interface{}, targetStruct interface{}) {
	structToStructSetDefault(sourceStruct, targetStruct, true)
}

// structToStructSetDefault 结构体转结构体
func structToStructSetDefault(sourceStruct interface{}, targetStruct interface{}, isSetDefault bool) {
	srcValue := reflect.ValueOf(sourceStruct)
	srcType := srcValue.Type()
	dstValue := reflect.ValueOf(targetStruct)

	dstValueE := dstValue.Elem()
	dstType := dstValueE.Type()

	for i := 0; i < srcValue.NumField(); i++ {
		sourceVF := srcValue.Field(i)
		sourceTF := srcType.Field(i)

		name := sourceTF.Name
		targetTF, ok := dstType.FieldByName(name)
		if !ok {
			continue
		}
		targetVF := dstValueE.FieldByName(name)
		if sourceVF.Type() != targetVF.Type() {
			continue
		}
		switch targetTF.Type.Kind() {
		case reflect.Struct:
			if targetTF.Type.Name() != "Time" {
				structToStructSetDefault(sourceVF.Interface(), targetVF.Addr().Interface(), isSetDefault)
			} else {
				reflectValue(sourceVF, targetTF, targetVF, isSetDefault)
			}
		case reflect.Slice:
			targetTE := targetTF.Type.Elem()
			elemDst := reflect.New(targetTE).Elem()
			sliceType := reflect.SliceOf(targetTE)
			fieldSlice := reflectSliceValue(sliceType, elemDst, sourceVF, isSetDefault)
			reflectValue(fieldSlice, targetTF, targetVF, isSetDefault)
		case reflect.Map:
			mapType := sourceTF.Type
			mapValue := reflect.MakeMap(mapType)
			reflectMapValue(sourceTF, sourceVF, targetTF, mapValue, isSetDefault)
			reflectValue(mapValue, targetTF, targetVF, isSetDefault)
		default:
			reflectValue(sourceVF, targetTF, targetVF, isSetDefault)
		}
	}
}

// reflectValue 设置字段值
func reflectValue(sourceV reflect.Value, targetT reflect.StructField, targetV reflect.Value, isSetDefault bool) {
	if sourceV.Type() != targetV.Type() || !sourceV.IsValid() {
		return
	}
	targetV.Set(sourceV)
	if isSetDefault {
		setStructDefaultValue(targetT, targetV)
	}
	return
}

// reflectSliceValue 设置切片默认值
func reflectSliceValue(sliceType reflect.Type, elemDst, sourceVF reflect.Value, isSetDefault bool) reflect.Value {
	fieldSlice := reflect.MakeSlice(sliceType, sourceVF.Len(), sourceVF.Len())
	for j := 0; j < sourceVF.Len(); j++ {
		elem := sourceVF.Index(j)
		switch elem.Kind() {
		case reflect.Struct:
			structToStructSetDefault(elem.Interface(), elemDst.Addr().Interface(), isSetDefault)
		case reflect.Slice:
			// 多维切片
			elemDst = elem
		default:
			elemDst = elem
		}
		fieldSlice.Index(j).Set(elemDst)
	}
	return fieldSlice
}

// reflectMapValue 设置map默认值
func reflectMapValue(sourceTF reflect.StructField, sourceVF reflect.Value, targetTF reflect.StructField, targetVF reflect.Value, isSetDefault bool) {
	for _, key := range sourceVF.MapKeys() {
		keyValue := key.Interface()
		valueK := sourceVF.MapIndex(key)
		switch valueK.Kind() {
		case reflect.Struct:
			elemDst := reflect.New(targetTF.Type.Elem()).Elem()
			structToStructSetDefault(valueK.Interface(), elemDst.Addr().Interface(), isSetDefault)
			targetVF.SetMapIndex(reflect.ValueOf(keyValue), elemDst)

		case reflect.Slice:
			elemDst := reflect.New(valueK.Type().Elem()).Elem()
			sliceType := targetTF.Type.Elem()
			sliceValue := sourceVF.MapIndex(key)
			fieldSlice := reflectSliceValue(sliceType, elemDst, sliceValue, isSetDefault)
			targetVF.SetMapIndex(reflect.ValueOf(keyValue), fieldSlice)

		case reflect.Map:
			// 多维map
			targetVF.SetMapIndex(reflect.ValueOf(keyValue), valueK)
		default:
			targetVF.SetMapIndex(reflect.ValueOf(keyValue), valueK)
		}
	}
}

// SetStructDefaultValue 设置default字段
func setStructDefaultValue(targetT reflect.StructField, targetV reflect.Value) {
	defaultValue := targetT.Tag.Get("default")
	if defaultValue == "" {
		return
	}
	currentV := targetV.Interface()
	switch targetT.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		defaultInt, _ := strconv.Atoi(defaultValue)
		setReflectDefaultInt(currentV, defaultInt, targetV)
	case reflect.Float32, reflect.Float64:
		defaultFloat, _ := strconv.ParseFloat(defaultValue, 64)
		setReflectDefaultFloat(currentV, defaultFloat, targetV)
	case reflect.String:
		defaultString := defaultValue
		setReflectDefaultString(currentV, defaultString, targetV)
	case reflect.Bool:
		defaultVal, _ := strconv.ParseBool(defaultValue)
		setReflectDefaultBool(currentV, defaultVal, targetV)
	default:
	}
}

// setReflectDefaultInt 设置int的default
func setReflectDefaultInt(currentValue interface{}, defaultValue int, targetV reflect.Value) {
	currentInt, ok := currentValue.(int)
	if !ok || currentInt != 0 {
		return
	}
	targetV.SetInt(int64(defaultValue))
}

// setReflectDefaultFloat 设置float的default
func setReflectDefaultFloat(currentValue interface{}, defaultValue float64, targetV reflect.Value) {
	currentInt, ok := currentValue.(float64)
	if !ok || currentInt != 0 {
		return
	}
	targetV.SetFloat(defaultValue)
}

// setReflectDefaultString 设置string的default
func setReflectDefaultString(currentValue interface{}, defaultValue string, targetV reflect.Value) {
	currentInt, ok := currentValue.(string)
	if !ok || currentInt != "" {
		return
	}
	targetV.SetString(defaultValue)

}

// setReflectDefaultBool 设置bool的default
func setReflectDefaultBool(currentValue interface{}, defaultValue bool, targetV reflect.Value) {
	currentVal, ok := currentValue.(bool)
	if !ok || !currentVal {
		return
	}
	targetV.SetBool(defaultValue)
}

// structToMap 结构体转map
func structToMap(req interface{}) map[string]interface{} {
	reqValue := reflect.ValueOf(req)
	reqType := reqValue.Type()
	reqMap := make(map[string]interface{})
	for i := 0; i < reqType.NumField(); i++ {
		field := reqType.Field(i)
		fieldValue := reqValue.Field(i)
		fieldName := field.Name

		if fieldValue.Kind() == reflect.Slice {
			// 如果字段是切片类型，递归处理切片元素
			var sliceData []map[string]interface{}
			for j := 0; j < fieldValue.Len(); j++ {
				sliceElement := fieldValue.Index(j).Interface()
				sliceElementMap := structToMap(sliceElement)
				sliceData = append(sliceData, sliceElementMap)
			}
			reqMap[fieldName] = sliceData
		} else if fieldValue.Kind() == reflect.Struct {
			fieldMap := structToMap(fieldValue.Interface())
			reqMap[fieldName] = fieldMap
		} else {
			reqMap[fieldName] = fieldValue
		}
	}
	return reqMap
}
