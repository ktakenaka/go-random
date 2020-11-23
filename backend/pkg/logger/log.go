package logger

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
}

// SetLogLevel set level
func SetLogLevel(level string) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	logger.SetLevel(lv)
}

// Info info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Panic panic
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Fatal fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Debugf debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Panicf panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatalf fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// WithError with error
func WithError(err error) *logrus.Entry {
	return logger.WithError(err)
}

// WithField with a field
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// WithFields with fields
func WithFields(fields map[string]interface{}) *logrus.Entry {
	return logger.WithFields(logrus.Fields(fields))
}
