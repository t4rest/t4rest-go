package logger

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger .
type Logger struct {
	*zap.SugaredLogger
	FastLog *zap.Logger
}

const fieldsKey = "logger_fields"

// Fields .
type Fields map[string]string

// New .
func New(cfg Conf) *Logger {

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	atom := zap.NewAtomicLevel()
	log := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))
	atom.SetLevel(zap.InfoLevel)

	if cfg.LogLevel != "" {
		err := atom.UnmarshalText([]byte(strings.ToLower(cfg.LogLevel)))
		if err != nil {
			log.Fatal("Invalid log level")
		}
	}

	fastLog := log.With(zap.String("service", cfg.AppID))
	return &Logger{
		SugaredLogger: fastLog.Sugar(),
		FastLog:       fastLog,
	}
}

// ContextWithFields .
func (l Logger) ContextWithFields(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, fieldsKey, fields)
}

// WithContext .
func (l Logger) WithContext(ctx context.Context) *zap.SugaredLogger {
	log := l.SugaredLogger
	fields, ok := ctx.Value(fieldsKey).(Fields)
	if !ok {
		return log
	}

	for k, v := range fields {
		log = log.With(k, v)
	}

	return log
}

// Flush .
func (l Logger) Flush() {
	_ = l.Sync()
	_ = l.FastLog.Sync()
}
