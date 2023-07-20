package log

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

const defaultLevel = "debug"

var levels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
	"panic": zapcore.PanicLevel,
}

func whichlevel(ls string) zapcore.Level {
	if l, found := levels[strings.ToLower(ls)]; found {
		return l
	}
	// debug is the default level
	return levels["debug"]
}
