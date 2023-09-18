package utils

import (
	"reflect"
	"strconv"
)

func StructToStruct(sourceStruct interface{}, targetStruct interface{}) {
	source := structToMap(sourceStruct)
	sourceValue := reflect.ValueOf(sourceStruct)

	targetV := reflect.ValueOf(targetStruct)
	targetT := reflect.TypeOf(targetStruct)
	if targetV.Kind() != reflect.Ptr {
		return
	}
	targetV = targetV.Elem()
	targetT = targetT.Elem()
	for i := 0; i < targetV.NumField(); i++ {
		targetValue := targetV.Field(i)
		targetType := targetT.Field(i)

		fieldName := targetType.Name
		sourceRV := source[fieldName]

		if targetValue.Kind() == reflect.Struct {
			reflectStructValue(sourceRV, targetValue)
			continue
		} else if targetValue.Kind() == reflect.Slice {
			reflectSliceValue(sourceValue, targetType, targetV)
			continue
		} else {
			reflectValue(sourceRV, targetType, targetValue)
		}
	}
	return
}

// reflectValue 设置字段值
func reflectValue(source interface{}, targetT reflect.StructField, targetV reflect.Value) {

	var (
		sourceVal reflect.Value
		ok        bool
	)
	sourceVal, ok = source.(reflect.Value)
	if !ok || !sourceVal.IsValid() || targetV.Type() != sourceVal.Type() {
		return
	}

	//if targetVal.Type() != sourceVal.Type() {
	//	if sourceVal.Type().String() != "time.Time" {
	//		return false
	//	}
	//	if sourceVal.Type() == reflect.TypeOf(time.Time{}) {
	//		timeValue := sourceVal.Interface().(time.Time)
	//		sourceVal = reflect.ValueOf(timeValue.Format(TIME_FORMAT))
	//	}
	//}
	targetV.Set(sourceVal)
	SetStructDefaultValue(targetT, targetV)
	return
}

// SetStructDefaultValue 设置default字段
func SetStructDefaultValue(targetT reflect.StructField, targetV reflect.Value) {
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

func reflectSliceValue(sourceValue reflect.Value, targetType reflect.StructField, targetVal reflect.Value) {
	fieldName := targetType.Name
	sourceFieldValue := sourceValue.FieldByName(fieldName)
	dataFieldValue := targetVal.FieldByName(fieldName)
	if dataFieldValue.CanSet() {
		// Copy the value from source field to data field
		dataFieldValue.Set(sourceFieldValue)
	}
}

// reflectStructValue 设置结构体的字段值
func reflectStructValue(source interface{}, targetVal reflect.Value) {
	sourceV := source.(map[string]interface{})
	targetT := targetVal.Type()
	for j := 0; j < targetVal.NumField(); j++ {
		targetValue := targetVal.Field(j)
		targetType := targetT.Field(j)
		sourceVal := sourceV[targetType.Name]
		if targetValue.Kind() == reflect.Struct {
			reflectStructValue(sourceVal, targetValue)
		} else if targetValue.Kind() == reflect.Slice {

		} else {
			reflectValue(sourceVal, targetType, targetValue)
		}
	}
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
