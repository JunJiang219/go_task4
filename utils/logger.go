package utils

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	BlueColor   = "\033[34m"
	YellowColor = "\033[33m"
	GreenColor  = "\033[32m"
	RedColor    = "\033[31m"
	ResetColor  = "\033[0m"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func myLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	str := l.CapitalString()
	switch l {
	case zapcore.InfoLevel:
		enc.AppendString(BlueColor + str + ResetColor)
	case zapcore.WarnLevel:
		enc.AppendString(GreenColor + str + ResetColor)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(RedColor + str + ResetColor)
	default:
		enc.AppendString(str)
	}
}

func initLogger() {
	// 创建日志目录
	os.MkdirAll("logs", 0755)

	consoleCfg := zap.NewDevelopmentEncoderConfig()
	consoleCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") //zapcore.ISO8601TimeEncoder
	consoleCfg.EncodeLevel = myLevelEncoder
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleCfg),
		os.Stdout,
		zapcore.DebugLevel,
	)

	fileCfg := zap.NewProductionEncoderConfig()
	fileCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") //zapcore.ISO8601TimeEncoder
	file, _ := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileCfg),
		file,
		zapcore.InfoLevel,
	)

	core := zapcore.NewTee(consoleCore, fileCore)
	logger = zap.New(core, zap.AddCaller())
}

func GetLogger() *zap.Logger {
	once.Do(initLogger)
	return logger
}
