package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"strings"
)

func init() {
	var (
		logPath string
		level   zapcore.Level
	)

	env := os.Getenv("ENV")
	switch env {
	case "dev":
		level = zap.DebugLevel
	case "prod":
		level = zap.InfoLevel
	default:
		level = zap.InfoLevel
	}

	if env == "prod" {
		logPath = "/var/log/app.log"
	} else {
		logPath = "./app.log"
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		}), zapcore.AddSync(os.Stdout)),
		level,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

// Info logs a message at the Info level.
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Error logs a message at the Error level.
func Error(msg string, err error, fields ...zap.Field) {
	pc, file, line, ok := runtime.Caller(1)
	funcName := ""
	if ok {
		parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		funcName = parts[len(parts)-1]
	}
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	fields = append(fields, zap.String("caller", fmt.Sprintf("%s:%d:%s", file, line, funcName)))
	Logger.Error(msg, fields...)
}
