package utils

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once sync.Once
)

func NewLog() *zap.SugaredLogger {
	var log *zap.SugaredLogger
	// 只执行一次的函数
	once.Do(func() {
		// 定义日志编码器的配置
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

		// 使用 lumberjack 进行日志轮替
		fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./logs/myapp.log", // 日志文件位置
			MaxSize:    2,                  // 日志文件最大大小 (MB)
			MaxBackups: 5,                  // 最多保留的备份数
			MaxAge:     28,                 // 日志文件的最大保存天数
			Compress:   false,              // 是否压缩备份的日志文件
		})

		// 设置日志输出到控制台
		consoleWriteSyncer := zapcore.AddSync(os.Stdout)

		// 定义一个zapcore.Core用于写入日志文件
		fileCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			fileWriteSyncer,
			zap.InfoLevel,
		)

		// 定义一个zapcore.Core用于写入控制台
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			consoleWriteSyncer,
			zap.InfoLevel,
		)

		// 将两个zapcore.Core组合在一起
		core := zapcore.NewTee(fileCore, consoleCore)

		// 初始化一个zap.Logger 并转换为 SugaredLogger
		logger := zap.New(core, zap.AddCaller())
		log = logger.Sugar()
	})
	return log
}
