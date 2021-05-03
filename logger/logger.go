package logger

import (
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

// Flush .
func (l Logger) Flush() {
	_ = l.Sync()
	_ = l.FastLog.Sync()
}
