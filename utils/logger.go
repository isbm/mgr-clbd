package utils

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

func GetTextLogger(level logrus.Level, out *os.File) *logrus.Logger {
	formatter := new(nested.Formatter)
	formatter.HideKeys = true
	formatter.FieldsOrder = []string{"component", "category"}
	formatter.ShowFullLevel = true

	logger := logrus.New()

	if out == nil {
		logger.Out = os.Stderr
		formatter.NoColors = false
		formatter.NoFieldsColors = false
	} else {
		logger.Out = out
		formatter.NoColors = true
		formatter.NoFieldsColors = true
	}

	logger.Level = level
	logger.SetFormatter(formatter)

	return logger
}

type GormLogger struct {
	lgr *logrus.Logger
}

func NewGormLogger(lg *logrus.Logger) *GormLogger {
	lgr := new(GormLogger)
	lgr.SetLogger(lg)
	return lgr
}

func (lgr *GormLogger) SetLogger(lg *logrus.Logger) *GormLogger {
	if lg == nil {
		lg = GetTextLogger(logrus.DebugLevel, nil)
	}
	lgr.lgr = lg
	return lgr
}

func (lgr *GormLogger) Print(v ...interface{}) {
	var data interface{}
	if v[0] == "sql" {
		data = v[3]
	} else if v[0] == "log" {
		data = v[2]
	}
	lgr.lgr.Debugln(data)
}
