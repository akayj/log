package log

import (
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewFileWriteSyncer(path string) zapcore.WriteSyncer {
	ws := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	return zapcore.AddSync(ws)
}

func NewConsoleWriteSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}
