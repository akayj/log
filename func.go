package log

import "go.uber.org/zap"

var slogger *zap.SugaredLogger

var (
	Print  = Debug  // Print identicals to Debug
	Printf = Debugf // Printf just like printf
	Printw = Debugw // Printw prints with structured kvargs
)

func whichlogger() *zap.SugaredLogger {
	if slogger == nil {
		InitDefault()
	}
	return slogger
}

func Debug(args ...interface{}) {
	whichlogger().Debug(args...)
}

func Debugf(msg string, args ...interface{}) {
	whichlogger().Debugf(msg, args...)
}

func Debugw(msg string, kvargs ...interface{}) {
	whichlogger().Debugw(msg, kvargs...)
}

func Info(args ...interface{}) {
	whichlogger().Info(args...)
}

func Infof(msg string, args ...interface{}) {
	whichlogger().Infof(msg, args...)
}

func Infow(msg string, kvargs ...interface{}) {
	whichlogger().Infow(msg, kvargs...)
}

func Warn(args ...interface{}) {
	whichlogger().Warn(args...)
}

func Warnf(msg string, args ...interface{}) {
	whichlogger().Warnf(msg, args...)
}

func Warnw(msg string, kvargs ...interface{}) {
	whichlogger().Warnw(msg, kvargs...)
}

func Error(args ...interface{}) {
	whichlogger().Error(args...)
}

func Errorf(msg string, args ...interface{}) {
	whichlogger().Errorf(msg, args...)
}

func Errorw(msg string, kvargs ...interface{}) {
	whichlogger().Errorw(msg, kvargs...)
}

func Fatal(args ...interface{}) {
	whichlogger().Fatal(args...)
}

func Fatalf(msg string, args ...interface{}) {
	whichlogger().Fatalf(msg, args...)
}

func Fatalw(msg string, kvargs ...interface{}) {
	whichlogger().Fatalw(msg, kvargs...)
}

func Panic(args ...interface{}) {
	whichlogger().Panic(args...)
}

func Panicf(msg string, args ...interface{}) {
	whichlogger().Panicf(msg, args...)
}

func Panicw(msg string, kvargs ...interface{}) {
	whichlogger().Panicw(msg, kvargs...)
}
