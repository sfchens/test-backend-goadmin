package utils

import (
	"errors"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type FileName struct {
	Name    string
	ExtName string
}

func GetFileName(filename string) (nameTmp FileName) {
	strArr := strings.Split(filename, ".")
	name := strings.Join(strArr[:len(strArr)-1], ".")
	nameTmp.Name = name
	nameTmp.ExtName = strArr[len(strArr)-1]
	return
}
