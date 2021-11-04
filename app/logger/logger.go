package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wrap *zap.Logger.
type Logger struct {
	*zap.Logger
}

var logger *zap.Logger

// SetLogger set logger variable.
func SetLogger() error {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.InfoLevel)

	config := zap.Config{
		Level:             level,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			LevelKey:         "level",
			TimeKey:          "time",
			NameKey:          "name",
			CallerKey:        "caller",
			FunctionKey:      "func",
			StacktraceKey:    "trace",
			LineEnding:       "",
			EncodeLevel:      zapcore.CapitalLevelEncoder,
			EncodeTime:       zapcore.RFC3339TimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			EncodeName:       zapcore.FullNameEncoder,
			ConsoleSeparator: "space",
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{},
	}

	var err error

	logger, err = config.Build()

	return err
}

// CloseLogger close the logger.
func CloseLogger() {
	if err := logger.Sync(); err != nil {
		log.Printf("fail to close logger (error = %v)\n", err)
	}
}

// NewLogger creates *Logger.
func NewLogger(name string) *Logger {
	return &Logger{logger.Named(name)}
}

// GetRawLogger returns *zap.Logger.
func (l *Logger) GetRawLogger() *zap.Logger {
	return l.Logger
}
