package easy_validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var easyValidator = easyValidate{}

type easyValidate struct {
	Validator *validator.Validate
}

func EasyValidator() easyValidate {
	return easyValidator
}

func InitPlayground() {
	validatorObj := validator.New()
	validatorObj.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("msg")
		if label == "" {
			return field.Name
		}
		return label
	})
	easyValidator.Validator = validatorObj
}

func (m easyValidate) Validate(s interface{}) (err error) {
	err = m.Validator.Struct(s)

	if err == nil {
		return
	}
	err = m.dealErrMsg(err)
	return
}

func (m easyValidate) dealErrMsg(err error) (errMsg error) {
	errMsg = err
	var (
		isExist bool
		field   string
	)

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return
	}
	for _, val := range err.(validator.ValidationErrors) {
		field = val.StructField()
		arr := strings.Split(val.Namespace(), ".")
		if len(arr) <= 0 {
			break
		}
		arr = strings.Split(arr[1], "|")
		if len(arr) <= 0 {
			break
		}
		for _, v := range arr {
			errMsgArr := strings.Split(v, ":")
			if errMsgArr[0] == val.Tag() && errMsgArr[1] != "" {
				errMsg = errors.New(errMsgArr[1])
				isExist = true
				break
			}
		}
	}

	if !isExist {
		errMsg = errors.New(fmt.Sprintf("请求参数 %s 校验不通过", field))
	}
	return
}
