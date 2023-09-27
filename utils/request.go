package utils

import (
	"csf/library/easy_config"
	"csf/library/easy_validator"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"regexp"
	"strconv"
)

func BindParams(ctx *gin.Context, datas ...interface{}) (err error) {
	for _, data := range datas {
		val := reflect.ValueOf(data)
		if val.Kind() != reflect.Ptr {
			err = errors.New("绑定参数传值类型异常")
			return
		}
		switch ctx.Request.Method {
		case "POST":
			contentType := ctx.Request.Header.Get("Content-Type")
			if regexp.MustCompile("application/json").MatchString(contentType) {
				err = ctx.ShouldBindJSON(data)
			} else {
				err = ctx.ShouldBind(data)
			}
		default:
			err = ctx.ShouldBindQuery(data)
		}
		if isSliceOrArray(data) {
			//err = CheckArr(data)
		} else {
			// 设置默认值
			err = SetDefault(data)
			if err != nil {
				return
			}
			// 校验参数是否合法
			err = easy_validator.NewValidator().Validate(data)
			if err != nil {
				return
			}
		}
	}
	return
}

func SetDefault(i interface{}) (err error) {
	typeOf := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)
	for i := 0; i < typeOf.Elem().NumField(); i++ {
		if valueOf.Elem().Field(i).IsZero() {
			def := typeOf.Elem().Field(i).Tag.Get("default")
			if def != "" {
				switch typeOf.Elem().Field(i).Type.String() {
				case "int":
					result, _ := strconv.Atoi(def)
					valueOf.Elem().Field(i).SetInt(int64(result))
				case "uint":
					result, _ := strconv.ParseUint(def, 10, 64)
					valueOf.Elem().Field(i).SetUint(result)
				case "string":
					valueOf.Elem().Field(i).SetString(def)
				}
			}
		}
	}
	_, err = json.Marshal(i)
	return
}

func GetRequestCurl(ctx *gin.Context, params interface{}) (reqUrl string) {
	path := ctx.Request.URL.Path
	query := ctx.Request.URL.RawQuery
	method := ctx.Request.Method

	url := fmt.Sprintf("%v:%d%v",
		easy_config.Config.App.BaseUrl,
		easy_config.Config.App.Port,
		path,
	)

	switch method {
	case "GET":
		url = fmt.Sprintf("'%v?%v'", url, query)
	case "POST":
		url = fmt.Sprintf("'%v' -H 'Content-Type: application/json' -d '%v'", url, ToJson(params))
	}
	reqUrl = fmt.Sprintf("curl -X %v %v", method, url)

	return
}

func isSliceOrArray(ptr interface{}) bool {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		return false
	}
	elemKind := rv.Elem().Kind()
	return elemKind == reflect.Slice || elemKind == reflect.Array
}
