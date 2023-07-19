package utils

import "time"

func NowDate() string {
	tpl := "2006-01-02"
	return time.Now().Format(tpl)
}
