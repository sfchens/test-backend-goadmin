package utils

import (
	"strings"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

func NowDate() string {
	tpl := "2006-01-02"
	return time.Now().Format(tpl)
}

func TimeNow() string {
	return time.Now().Format(TIME_FORMAT)
}

func SplitTimeStr(timeStr string, typ string) (timeStrTmp string) {
	var timeArr []string
	switch typ {
	case "CST":
		timeArr = strings.Split(timeStr, " ")
	default:
	}
	if len(timeArr) >= 2 {
		timeStrTmp = timeArr[0] + " " + timeArr[1]
	} else if len(timeArr) == 1 {
		timeStrTmp = timeArr[0]
	}
	return
}
