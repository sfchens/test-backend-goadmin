package easy_logger

import (
	"csf/library/global"
	"time"
)

const LogFileApp = "app"                   // 常规日志
const LogFileAdminKey = global.ModuleAdmin // 常规日志
const LogFileH5Key = global.ModuleH5       // 常规日志
const LogFileLoginKey = "login"            // 常规日志

var LogArr = []string{
	LogFileApp,
	LogFileAdminKey,
	LogFileH5Key,
	LogFileLoginKey,
}

type LoggerConfig struct {
	LogObjKey string
	LogMysql  bool //是否写入 mysql 数据库
	LogZap    bool // 是否写入 zap 日志中
	LogZinc   bool // 是否发送到 ZincSearch 服务器
}

type logData struct {
	TraceId      string                 `json:"trace_id" description:"trace_id"`
	Ip           string                 `json:"id" description:"IP"`
	UserAgent    string                 `json:"user-agent" description:"浏览器代理"`
	Duration     time.Duration          `json:"duration" description:"运行时长"`
	RequestUrl   string                 `json:"request_url" description:"请求地址"`
	Params       map[string]interface{} `json:"params" description:"参数"`
	Status       int                    `json:"status" description:"状态"`
	ResponseData map[string]interface{} `json:"response_data" description:"响应数据"`
}
