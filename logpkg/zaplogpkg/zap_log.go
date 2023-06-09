package zaplogpkg

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
	"time"
)

var zapLogger *zap.Logger
var zapOnce sync.Once

type LogConfig struct {
	File       string
	MaxSize    int  // MB
	MaxBackups int  // 保留旧文件的最大个数
	MaxAge     int  // days
	Compress   bool // 是否压缩 / 归档旧文件
}

func LogInit(config *LogConfig) {
	zapOnce.Do(func() {
		hook := lumberjack.Logger{
			Filename:   "./log",
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     7,
			Compress:   config.Compress,
		}
		if len(config.File) > 0 {
			hook.Filename = config.File
		}
		if config.MaxSize > 0 {
			hook.MaxSize = config.MaxSize
		}
		if config.MaxBackups > 0 {
			hook.MaxBackups = config.MaxBackups
		}
		if config.MaxAge > 0 {
			hook.MaxAge = config.MaxAge
		}

		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}

		core := zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(&hook),
				zap.InfoLevel,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.Lock(os.Stderr),
				zap.ErrorLevel,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.Lock(os.Stdout),
				zap.DebugLevel,
			),
		)

		zapLogger = zap.New(core)
	})
}

func Info(tag, msg string, err error) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Info("["+tag+"] "+msg, zap.Error(err))
	return nil
}

func Error(tag, msg string, err error) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Error("["+tag+"] "+msg, zap.Error(err))
	return nil
}

func Fatal(tag, msg string, err error) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Fatal("["+tag+"] "+msg, zap.Error(err))
	return nil
}

func Debug(tag, msg string, err error) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Debug("["+tag+"] "+msg, zap.Error(err))
	return nil
}

func Warn(tag, msg string, err error) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Warn("["+tag+"] "+msg, zap.Error(err))
	return nil
}

func Panic(tag, msg string, err error) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Panic("["+tag+"] "+msg, zap.Error(err))
	return nil
}

func LogSync() error {
	return zapLogger.Sync()
}
