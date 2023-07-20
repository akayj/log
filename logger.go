package log

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultOptions = []zap.Option{
	zap.AddCaller(),
	zap.AddCallerSkip(1),
}

func newLogger(enc zapcore.Encoder, ws zapcore.WriteSyncer, level string, opts ...zap.Option) *zap.Logger {
	core := zapcore.NewCore(enc, ws, whichlevel(level))
	return zap.New(core, opts...)
}

func newFileLogger(path, level string, opts ...zap.Option) *zap.Logger {
	return newLogger(NewJSONEncoder(), NewFileWriteSyncer(path), level, opts...)
}

func newConsoleLogger(level string, opts ...zap.Option) *zap.Logger {
	return newLogger(NewConsoleEncoder(), NewConsoleWriteSyncer(), level, opts...)
}

func InitDefault() {
	InitWithLevel(defaultLevel)
}

func InitWithLevel(level string) {
	logger := newConsoleLogger(level, defaultOptions...)
	slogger = logger.Sugar()
}

func InitWithFields(fieldKvs ...string) {
	InitWithLevelFields(defaultLevel, fieldKvs...)
}

func InitWithLevelFields(level string, fieldKvs ...string) {
	if l := NewWithLevelFields(level, fieldKvs...); l != nil {
		slogger = l.Sugar()
	}
}

func InitFilelogger(path string) {
	logger := NewFileLoggerWithLevel(path, defaultLevel)
	slogger = logger.Sugar()
}

func InitFileLoggerWithLevel(path, level string) {
	l := NewFileLoggerWithLevel(path, level)
	slogger = l.Sugar()
}

func InitFileLoggerWithFields(path string, fieldKvs ...string) {
	l := NewFileLoggerWithFields(path, fieldKvs...)
	slogger = l.Sugar()
}

func InitFileLoggerWithLevelFields(path, level string, fieldKvs ...string) {
	l := NewFileLoggerWithLevelFields(path, level, fieldKvs...)
	slogger = l.Sugar()
}

func New(level ...string) *zap.Logger {
	ls := defaultLevel
	if len(level) > 0 {
		ls = level[0]
	}
	return newConsoleLogger(ls, defaultOptions...)
}

func NewWithFields(fieldKvs ...string) *zap.Logger {
	return NewWithLevelFields(defaultLevel, fieldKvs...)
}

func NewWithLevelFields(level string, filedKvs ...string) *zap.Logger {
	opt, err := FieldsOption(filedKvs...)
	if err != nil {
		return nil
	}

	return NewWithOptions(level, opt)
}

func NewWithOptions(level string, opts ...zap.Option) *zap.Logger {
	l := New()
	return WithOptions(l, opts...)
}

func NewFileLogger(path string) *zap.Logger {
	return NewFileLoggerWithLevelFields(path, defaultLevel)
}

func WithOptions(l *zap.Logger, opts ...zap.Option) *zap.Logger {
	return l.WithOptions(opts...)
}

func NewFileLoggerWithLevel(path, level string) *zap.Logger {
	return newFileLogger(path, level)
}

func NewFileLoggerWithFields(path string, fieldKvs ...string) *zap.Logger {
	return NewFileLoggerWithLevelFields(path, defaultLevel, fieldKvs...)
}

func NewFileLoggerWithLevelFields(path, level string, fieldKvs ...string) *zap.Logger {
	fieldsOpt, err := FieldsOption(fieldKvs...)
	if err != nil {
		return nil
	}

	return newFileLogger(path, level, fieldsOpt)
}

func FieldsOption(fieldKvs ...string) (zap.Option, error) {
	if len(fieldKvs)%2 != 0 {
		return nil, errors.New("invalid fieldKvs")
	}

	fields := make([]zap.Field, 0)
	for i := 0; i < len(fieldKvs)/2; i++ {
		fields = append(fields, zap.String(fieldKvs[i], fieldKvs[i+1]))
	}

	return zap.Fields(fields...), nil
}
