package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log holds an instance of zap logger.
var log *zap.Logger

// init will be called when this package is imported.
func init() {
	logConf := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	log, err = logConf.Build()
	if err != nil {
		panic(err)
	}
}

// Info logs a message at InfoLevel.
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

// Warn logs a message at WarnLevel.
func Warn(msg string, tags ...zap.Field) {
	log.Warn(msg, tags...)
	log.Sync()
}

// Error logs a message at ErrorLevel.
func Error(err error, tags ...zap.Field) {
	log.Error(err.Error(), tags...)
	log.Sync()
}

// Fatal logs a message at FatalLevel.
func Fatal(err error, tags ...zap.Field) {
	log.Fatal(err.Error(), tags...)
	log.Sync()
}
