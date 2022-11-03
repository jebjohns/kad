package logging

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Audit(auditType, operation, status, user, format string, args ...interface{})
}

type Logging struct {
}

func (l *Logging) Infof(format string, args ...interface{}) {
	fileDetails := l.getFileDetails()
	logrus.WithFields(
		logrus.Fields{
			"caller": fileDetails,
		},
	).Infof(format, args...)
}

func (l *Logging) Debugf(format string, args ...interface{}) {
	fileDetails := l.getFileDetails()
	logrus.WithFields(
		logrus.Fields{
			"caller": fileDetails,
		},
	).Debugf(format, args...)
}

func (l *Logging) Errorf(format string, args ...interface{}) {
	fileDetails := l.getFileDetails()
	logrus.WithFields(
		logrus.Fields{
			"caller": fileDetails,
		},
	).Errorf(format, args...)
}

func (l *Logging) Fatalf(format string, args ...interface{}) {
	fileDetails := l.getFileDetails()
	logrus.WithFields(
		logrus.Fields{
			"caller": fileDetails,
		},
	).Fatalf(format, args...)
}

func (l *Logging) Audit(auditType, operation, status, user, format string, args ...interface{}) {
	fileDetails := l.getFileDetails()
	logrus.WithFields(
		logrus.Fields{
			"type":      auditType,
			"operation": operation,
			"status":    status,
			"user":      user,
			"caller":    fileDetails,
		},
	).Infof(format, args...)
}

func (*Logging) getFileDetails() string {
	pc, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%v:%s", file, line, path.Base(runtime.FuncForPC(pc).Name()))
}

func NewLogger() Logger {
	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	return &Logging{}
}