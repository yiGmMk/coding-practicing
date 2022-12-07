package zlog

import (
	"os"
	"path"
	"runtime"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func ZlogInit(path string, rotate bool, options ...zap.Option) {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	var fileWriteSyncer zapcore.WriteSyncer
	if !rotate {
		file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		fileWriteSyncer = zapcore.AddSync(file)
	} else {
		fileWriteSyncer = getFileLogWriter(path)
	}

	core := zapcore.NewTee(
		// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
	)
	logger = zap.New(core)
}

// 日志切割
func getFileLogWriter(path string) (writeSyncer zapcore.WriteSyncer) {
	// 使用 lumberjack 实现 log rotate
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    100, // 单个文件最大100M
		MaxBackups: 60,  // 多于 60 个日志文件后，清理较旧的日志
		MaxAge:     1,   // 一天一切割
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {

	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
