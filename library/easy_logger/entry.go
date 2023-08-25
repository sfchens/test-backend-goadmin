package easy_logger

import (
	"csf/library/easy_config"
	"csf/library/global"
	"csf/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Logger logger

type logger struct {
	CustomZap  _customZap
	CurrentZap *zap.Logger
}

func InitLogger() {
	// 初始化zap日志管理
	Logger = logger{}
	Logger.CustomZap = initCustomZap()
}

// ZapLog zap.Logger对象
func ZapLog(fileNames ...string) (loggerT *zap.Logger) {
	log := Logger.CustomZap.ZapLogger

	if len(fileNames) > 0 {
		if log[fileNames[0]] == nil {
			loggerT = zapLogWith(fileNames[0])
		} else {
			loggerT = log[fileNames[0]]
		}
	} else {
		loggerT = log[LogFileAppKey]
	}
	// 显示抛出异常行数
	if easy_config.Config.Zap.ShowLine {
		loggerT = loggerT.WithOptions(zap.AddCaller())
	}
	Logger.CurrentZap = loggerT
	return loggerT
}

// OperateLogger 操作日志
func OperateLogger(ctx *gin.Context, logObjKey string) {
	logDataObj := getRequestData(ctx)
	msg := GetLogTemplate(ctx, nil)

	zapLog := ZapLog(logObjKey)
	// 记录操作信息
	if logDataObj.Status != http.StatusOK {
		zapLog.Error(msg)
	}
	zapLog.Info(msg)
}

// getDefaultLog 获取默认日志名
func getDefaultLog() []string {
	logArr := []string{
		LogFileAppKey,
	}
	return logArr
}

// getRequestData 获取请求参数
func getRequestData(ctx *gin.Context) logData {
	traceId := ctx.GetString(global.TraceIdKey)
	// 请求参数
	params, _ := utils.GetRequestParams(ctx)

	// 结果信息
	responseWriter := utils.ResponseWriter
	duration := time.Since(responseWriter.StartTime)
	var resp map[string]interface{}
	json.Unmarshal([]byte(responseWriter.Body.String()), &resp)
	tmpData := logData{
		TraceId:      traceId,
		Ip:           ctx.ClientIP(),
		Status:       ctx.Writer.Status(),
		UserAgent:    ctx.Request.UserAgent(),
		RequestUrl:   utils.GetCurl(ctx),
		Params:       params,
		Duration:     duration,
		ResponseData: resp,
	}
	return tmpData
}

// GetLogTemplate 日志模板
func GetLogTemplate(ctx *gin.Context, msg interface{}) string {
	logDataObj := getRequestData(ctx)
	tmp := fmt.Sprintf("[%v][%vmm][%v][%v]# ", logDataObj.Ip, logDataObj.Duration.Seconds(), logDataObj.TraceId, utils.GetUserName(ctx))
	if msg != nil {
		var msgT interface{}
		switch msg.(type) {
		case string, int, int8, int16, int32, int64, float64, float32:
			msgT = fmt.Sprintf("%v", msg)
		default:
			msgT = utils.ToJson(msg)
		}
		tmp = fmt.Sprintf("%v %v", tmp, msgT)
	} else {
		tmp = fmt.Sprintf("%v %v response: '%v'", tmp, utils.GetCurl(ctx), utils.ToJson(logDataObj.ResponseData))
	}
	return tmp
}

func (s *logger) Info(ctx *gin.Context, msg interface{}) {
	tmpMsg := GetLogTemplate(ctx, msg)
	s.CurrentZap.Info(tmpMsg)
}
func (s *logger) Warn(ctx *gin.Context, msg interface{}) {
	tmpMsg := GetLogTemplate(ctx, msg)
	s.CurrentZap.Warn(tmpMsg)
}
func (s *logger) Error(ctx *gin.Context, msg interface{}) {
	tmpMsg := GetLogTemplate(ctx, msg)
	s.CurrentZap.Error(tmpMsg)
}
func (s *logger) Panic(ctx *gin.Context, msg interface{}) {
	tmpMsg := GetLogTemplate(ctx, msg)
	s.CurrentZap.Panic(tmpMsg)
}
func (s *logger) Fatal(ctx *gin.Context, msg interface{}) {
	tmpMsg := GetLogTemplate(ctx, msg)
	s.CurrentZap.Fatal(tmpMsg)
}
func (s *logger) Debug(ctx *gin.Context, msg interface{}) {
	tmpMsg := GetLogTemplate(ctx, msg)
	s.CurrentZap.Debug(tmpMsg)
}
