package easy_logger

import (
	"csf/library/easy_config"
	"csf/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

var customZap _customZap

type _customZap struct {
	ZapLogger map[string]*zap.Logger
	fileName  string
	level     string
}

// InitCustomZap 初始化zap
func initCustomZap() _customZap {
	paths := easy_config.Viper.GetString("zap.path")
	if ok, _ := utils.FileExists(paths); !ok {
		_ = os.Mkdir(paths, os.ModePerm)
	}

	fileNamesTmp := easy_config.Viper.GetStringSlice("zap.log-file")
	if len(fileNamesTmp) <= 0 {
		fileNamesTmp = getDefaultLog()
	}
	for _, fileName := range fileNamesTmp {
		if len(customZap.ZapLogger) <= 0 {
			customZap = _customZap{
				ZapLogger: make(map[string]*zap.Logger),
			}
		}
		customZap.fileName = fileName

		cores := customZap.GetZapCores()

		loggerT := zap.New(zapcore.NewTee(cores...))
		customZap.ZapLogger[fileName] = loggerT
		customZap.level = easy_config.Viper.GetString("zap.level")
	}

	return customZap
}

// zapLogWith 根据目录创建日志
func zapLogWith(fileName string) *zap.Logger {
	customZap = _customZap{
		fileName: fileName,
		level:    easy_config.Viper.GetString("zap.level"),
	}
	cores := customZap.GetZapCores()
	loggerT := zap.New(zapcore.NewTee(cores...))
	return loggerT
}

// GetEncoder 获取 zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) GetEncoder() zapcore.Encoder {
	if easy_config.Viper.GetString("zap.format") == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",                                         // 日志内容对应的key名，此参数必须不为空，否则日志主体不处理
		LevelKey:       "level",                                           // 日志级别对应的key名
		TimeKey:        "time",                                            // 时间对应的key名
		NameKey:        "logger",                                          // logger名对应的key名
		CallerKey:      "caller",                                          // 调用者对应的key名
		StacktraceKey:  easy_config.Viper.GetString("zap.stacktrace-key"), // 栈追踪的key名
		LineEnding:     zapcore.DefaultLineEnding,                         // 行末输出格式
		EncodeLevel:    z.ZapEncodeLevel(),                                // 日志编码级别
		EncodeTime:     z.CustomTimeEncoder,                               // 日志时间解析
		EncodeDuration: zapcore.SecondsDurationEncoder,                    // 日志日期解析
		EncodeCaller:   zapcore.FullCallerEncoder,                         // 日志调用路径
	}
}

// GetEncoderCore 获取Encoder的 zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := z.GetWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *_customZap) ZapEncodeLevel() zapcore.LevelEncoder {
	level := easy_config.Viper.GetString("zap.encode-level")
	switch {
	case level == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case level == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case level == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case level == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// GetWriteSyncer 切割日志，指定文件
func (z *_customZap) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename:   z.GetPath(level),
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     easy_config.Viper.GetInt("zap.max-age"), // days
		Compress:   false,                                   //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}
	if easy_config.Viper.GetBool("zap.log-in-console") {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}

// GetPath 获取路径
func (z *_customZap) GetPath(level string) string {
	var pathArr = []string{
		easy_config.Viper.GetString("zap.path"),
		time.Now().Format("2006-01-02"),
	}
	if len(z.fileName) > 0 {
		pathArr = append(pathArr, z.fileName)
	}
	pathArr = append(pathArr, level+".log")
	pathStr := path.Join(pathArr...)
	return pathStr
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	var options = []string{
		easy_config.Viper.GetString("zap.prefix"),
		t.Format("2006-01-02 15:04:05.000"),
	}
	encoder.AppendString(strings.Join(options, " "))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}

	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// TransportLevel 根据字符串转化为 zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_customZap) TransportLevel() zapcore.Level {
	level := z.level
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// MapToFields map转Fields
func (z *_customZap) MapToFields(logData logData) (fields []zapcore.Field) {
	targetV := reflect.ValueOf(logData)
	if targetV.Kind() == reflect.Ptr {
		targetV = targetV.Elem()
	}
	for i := 0; i < targetV.NumField(); i++ {

		field := targetV.Field(i)
		targetT := targetV.Type()

		jsonName := targetT.Field(i).Tag.Get("json")

		fieldType := field.Interface()
		switch fieldType.(type) {
		case int:
			val, _ := (fieldType).(int)
			fields = append(fields, zap.Int(jsonName, val))
		case string:
			val, _ := (fieldType).(string)
			fields = append(fields, zap.String(jsonName, val))
		case time.Duration:
			val, _ := (fieldType).(time.Duration)
			fields = append(fields, zap.Duration(jsonName, val))
		case any:
			val, _ := (fieldType).(any)
			fields = append(fields, zap.Any(jsonName, val))
		default:
			val, _ := (fieldType).(any)
			fields = append(fields, zap.Any(jsonName, val))
			fmt.Printf("defalt: key：%v; value: %v\n", jsonName, val)
		}
	}
	return
}
