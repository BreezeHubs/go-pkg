package zaplogpkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type ZapLog struct {
	logger *zap.Logger
}

func LogNew(config *LogConfig) *ZapLog {
	z := &ZapLog{}

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
		//zapcore.NewCore(
		//	zapcore.NewJSONEncoder(encoderConfig),
		//	zapcore.Lock(os.Stdout),
		//	zap.DebugLevel,
		//),
	)

	z.logger = zap.New(core)
	return z
}

func (l *ZapLog) Info(tag, msg string, err error) {
	l.logger.Info("["+tag+"] "+msg, zap.Error(err))
}

func (l *ZapLog) Error(tag, msg string, err error) {
	l.logger.Error("["+tag+"] "+msg, zap.Error(err))
}

func (l *ZapLog) Fatal(tag, msg string, err error) {
	l.logger.Fatal("["+tag+"] "+msg, zap.Error(err))
}

func (l *ZapLog) Debug(tag, msg string, err error) {
	l.logger.Debug("["+tag+"] "+msg, zap.Error(err))
}

func (l *ZapLog) Warn(tag, msg string, err error) {
	l.logger.Warn("["+tag+"] "+msg, zap.Error(err))
}

func (l *ZapLog) Panic(tag, msg string, err error) {
	l.logger.Panic("["+tag+"] "+msg, zap.Error(err))
}

func (l *ZapLog) LogSync() error {
	return l.logger.Sync()
}
