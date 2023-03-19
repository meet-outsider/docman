package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 创建日志文件保存的目录，以及日志文件名
	logDir := "logs/" + time.Now().Format("2006-01-02")
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create log directory: %v", err)
	}
	infoLogFile := filepath.Join(logDir, "info.log")
	//errorLogFile := filepath.Join(logDir, "error.log")

	// 创建 lumberjack 实例，用于实现日志文件轮转
	lumberJackLogger := &lumberjack.Logger{
		Filename:   infoLogFile,
		MaxSize:    5, // 每个日志文件最大 5MB
		MaxBackups: 3, // 最多保留 3 个旧日志文件
		MaxAge:     7, // 旧日志文件最多保留 7 天
		Compress:   true,
	}

	// 创建 Zap 的配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger)),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	// 如果是开发环境，添加控制台输出
	if os.Getenv("ENV") == "dev" {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		)
	}

	// 创建 Zap 的 logger
	logger := zap.New(core)

	// 记录日志
	logger.Info("this is an info log")
	logger.Error("this is an error log")

	// 关闭 logger 和 lumberjack 实例
	if err := logger.Sync(); err != nil {
		fmt.Printf("failed to flush logs: %v", err)
	}
	if err := lumberJackLogger.Close(); err != nil {
		fmt.Printf("failed to close logger: %v", err)
	}
}
