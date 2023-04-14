package logpkg

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

func ZapLogInit(logFile string) {
	zapOnce.Do(func() {
		hook := lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    10, // 10MB
			MaxBackups: 10,
			MaxAge:     7, // 7 days
			Compress:   true,
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

func ZapLogInfo(lm *LogMessage) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Info(lm.Msg,
		zap.String("tag", lm.Tag),
		zap.Error(lm.Err),
	)
	return nil
}

func ZapLogError(lm *LogMessage) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Error(lm.Msg,
		zap.String("tag", lm.Tag),
		zap.Error(lm.Err),
	)
	return nil
}

func ZapLogFatal(lm *LogMessage) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Fatal(lm.Msg,
		zap.String("tag", lm.Tag),
		zap.Error(lm.Err),
	)
	return nil
}

func ZapLogDebug(lm *LogMessage) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Debug(lm.Msg,
		zap.String("tag", lm.Tag),
		zap.Error(lm.Err),
	)
	return nil
}

func ZapLogWarn(lm *LogMessage) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Warn(lm.Msg,
		zap.String("tag", lm.Tag),
		zap.Error(lm.Err),
	)
	return nil
}

func ZapLogPanic(lm *LogMessage) error {
	if zapLogger == nil {
		return errors.New("log 未初始化")
	}

	zapLogger.Panic(lm.Msg,
		zap.String("tag", lm.Tag),
		zap.Error(lm.Err),
	)
	return nil
}

func ZapLogSync() error {
	return zapLogger.Sync()
}
