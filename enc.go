package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewJSONEncoder() zapcore.Encoder {
	encconfig := zap.NewProductionEncoderConfig()

	encconfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encconfig.EncodeTime = zapcore.RFC3339TimeEncoder

	encconfig.TimeKey = "time"
	encconfig.CallerKey = "file"
	encconfig.MessageKey = "content"

	return zapcore.NewJSONEncoder(encconfig)
}

func NewConsoleEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()

	return zapcore.NewConsoleEncoder(cfg)
}
