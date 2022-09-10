package logger

import (
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/kmaguswira/coinbit/application/config"
	"github.com/sirupsen/logrus"
)

var logger LoggerStruct

type LoggerStruct struct {
	Log *logrus.Logger
}

func Init() {
	logger.Log = logrus.New()
	logger.Log.SetLevel(logrus.StandardLogger().Level)
	logger.Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})

	if config.GetConfig().Env != "development" {
		logger.Log.SetFormatter(&logrus.JSONFormatter{})
	}
}

func Log() LoggerStruct {
	return logger
}

func (t LoggerStruct) Info(msg ...string) {
	t.Log.WithFields(logrus.Fields{
		"caller": getCaller(1),
	}).Info(strings.Join(msg, "::"))
}

func (t LoggerStruct) Debug(msg ...string) {
	t.Log.WithFields(logrus.Fields{
		"caller": getCaller(1),
	}).Debug(strings.Join(msg, "::"))
}

func (t LoggerStruct) Warn(msg ...string) {
	t.Log.WithFields(logrus.Fields{
		"caller": getCaller(1),
	}).Warn(strings.Join(msg, "::"))
}

func (t LoggerStruct) Error(err error) {
	if err != nil {
		t.Log.WithFields(logrus.Fields{
			"caller": getCaller(1),
			"stack":  string(debug.Stack()),
		}).WithError(err).Error(err.Error())
	}
}

func getCaller(skip int) string {
	defaultSkip := 1
	if skip > 0 {
		defaultSkip += skip
	}

	pc, _, line, ok := runtime.Caller(defaultSkip)
	if !ok {
		return "unknown:0"
	}

	name := runtime.FuncForPC(pc).Name()
	lsi := strings.LastIndexByte(name, '/') + 1
	if lsi < 1 {
		lsi = 0
	}

	return name[lsi:] + ":" + strconv.Itoa(line)
}
