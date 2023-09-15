package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	defaultLogger = NewLogger()
	atomicLevel   = zap.NewAtomicLevel()
)

func NewLogger() *zap.Logger {
	var encCfg zapcore.EncoderConfig

	var encoder zapcore.Encoder

	encCfg = zap.NewProductionEncoderConfig()
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encCfg.LevelKey = zapcore.DebugLevel.String()

	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.TimeKey = "timestamp"

	encoder = zapcore.NewConsoleEncoder(encCfg)

	l := zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atomicLevel),
	)
	l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return l
}

func L(ctx context.Context) *zap.Logger {
	return LoggerFromContext(ctx)
}

func WithFields(ctx context.Context, fields ...zap.Field) *zap.Logger {
	return LoggerFromContext(ctx).With(fields...)
}

func WithError(ctx context.Context, err error) *zap.Logger {
	return LoggerFromContext(ctx).With(zap.Error(err))
}
