package log

import (
	conf "docman/config"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	consoleLogger *zap.Logger
	fileLogger    *zap.Logger
	debug         bool
)

func init() {
	debug = conf.Config.Server.Env != "dev"
	// 初始化控制台logger
	consoleEncoder := getEncoder(false)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	consoleLogger = zap.New(consoleCore)

	// 初始化文件logger
	fileEncoder := getFileEncoder(true)
	fileWriter := getLogWriter()
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), zapcore.DebugLevel)
	fileLogger = zap.New(fileCore, zap.AddCaller())
}

// 日志编码格式
func getEncoder(isFile bool) zapcore.Encoder {
	consoleSeparator := ""
	//if runtime.GOOS == "darwin" {
	//	consoleSeparator = " | "
	//} else {
	consoleSeparator = " \033[0m|\033[0m "
	//}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		//EncodeLevel:   zapcore.LowercaseColorLevelEncoder,
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			color := ""
			switch l {
			case zapcore.DebugLevel:
				color = "\033[36m" // Cyan
			case zapcore.InfoLevel:
				color = "\033[32m" // Green
			case zapcore.WarnLevel:
				color = "\033[33m" // Yellow
			case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
				color = "\033[31m" // Red
			}
			enc.AppendString(color + l.CapitalString() + "\033[0m")
		},
		LineEnding:       "",
		ConsoleSeparator: consoleSeparator,
		//EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		//	enc.AppendString(t.Format("2006-01-02 15:04:05"))
		//},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("\033[37m" + t.Format("2006-01-02 15:04:05") + "\033[0m")
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if isFile {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getFileEncoder(isFile bool) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		LineEnding:    "",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if isFile {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志写日文件配置
func getLogWriter() zapcore.WriteSyncer {
	logFilePath := "./logs/"
	logFileName := time.Now().Format("2006-01-02") + ".log"
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		if err != nil {
			if e := os.Mkdir(logFilePath, os.ModePerm); e != nil {
				Error(e.Error())
			}
		}
	}
	config := conf.Config.Logger
	fileWriter := &lumberjack.Logger{
		Filename:   filepath.Join(logFilePath, logFileName),
		MaxSize:    config.MaxSize,    // MB
		MaxAge:     config.MaxAge,     // days
		MaxBackups: config.MaxBackups, // files
		LocalTime:  true,
		Compress:   true,
	}
	return zapcore.AddSync(fileWriter)
}
func toJson(data ...any) string {
	var marshal, _ = json.Marshal(data)
	return string(marshal)
}
func Info(msg string, data ...any) {
	if debug {
		consoleLogger.Info(msg, zap.String("data", toJson(data)))
	}
	fileLogger.Info(msg, zap.String("data", toJson(data)))
}
func Error(msg string, data ...any) {
	if debug {
		consoleLogger.Error(msg, zap.String("data", toJson(data)))
	}
	fileLogger.Error(msg, zap.String("data", toJson(data)))
}
func Warn(msg string, data ...any) {
	if debug {
		consoleLogger.Warn(msg, zap.String("data", toJson(data)))
	}
	fileLogger.Warn(msg, zap.String("data", toJson(data)))
}
func Debug(msg string, data ...any) {
	if debug {
		consoleLogger.Debug(msg, zap.String("data", toJson(data)))
	}
	fileLogger.Debug(msg, zap.String("data", toJson(data)))
}
