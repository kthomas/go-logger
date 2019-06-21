package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	glogger "github.com/maratk/logger"
)

type Logger struct {
	console bool
	syslog  bool
	logger  *glogger.Logger
	prefix  string
	logPath *string
}

func (lg *Logger) configure() {
	if lg.logPath != nil {
		lf, err := os.OpenFile(*lg.logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			glogger.Fatalf("Failed to open log file: %v", err)
		}
		defer lf.Close()

		lg.logger = glogger.Init(lg.prefix, lg.console, lg.syslog, lf)
	} else {
		lg.logger = glogger.Init(lg.prefix, lg.console, lg.syslog, ioutil.Discard)
	}
	glogger.SetFlags(log.LstdFlags)

	var logPrefix = lg.prefix
	if len(lg.prefix) > 0 {
		logPrefix = fmt.Sprintf("%s ", logPrefix)
	}
}

func (lg *Logger) Clone() *Logger {
	return &Logger{
		console: lg.console,
		syslog:  lg.syslog,
		logger:  lg.logger,
		prefix:  lg.prefix,
		logPath: lg.logPath,
	}
}

func (lg *Logger) Critical(msg string) {
	lg.logger.Fatal(msg)
}

func (lg *Logger) Criticalf(msg string, v ...interface{}) {
	args := lg.transformParams(v)
	lg.logger.Fatalf(msg, args...)
}

func (lg *Logger) Debug(msg string) {
	lg.logger.Info(msg)
}

func (lg *Logger) Debugf(msg string, v ...interface{}) {
	args := lg.transformParams(v)
	lg.logger.Infof(msg, args...)
}

func (lg *Logger) Error(msg string) {
	lg.logger.Error(msg)
}

func (lg *Logger) Errorf(msg string, v ...interface{}) {
	args := lg.transformParams(v)
	lg.logger.Errorf(msg, args...)
}

func (lg *Logger) Info(msg string) {
	lg.logger.Info(msg)
}

func (lg *Logger) Infof(msg string, v ...interface{}) {
	args := lg.transformParams(v)
	lg.logger.Infof(msg, args...)
}

func (lg *Logger) LogOnError(err error, s string) bool {
	hasErr := false
	if err != nil {
		msg := fmt.Sprintf("Error: %s", err)
		if s != "" {
			msg = fmt.Sprintf("%s; %s", msg, s)
		}
		lg.Errorf(msg)
		hasErr = true
	}
	return hasErr
}

func (lg *Logger) Panicf(msg string, v ...interface{}) {
	args := lg.transformParams(v)
	lg.logger.Fatalf(msg, args...)
}

func (lg *Logger) PanicOnError(err error, s string) {
	if err != nil {
		msg := fmt.Sprintf("Error: %s", err)
		if s != "" {
			msg = fmt.Sprintf("%s; %s", msg, s)
		}
		lg.Panicf(msg)
	}
}

func (lg *Logger) Warning(msg string) {
	lg.logger.Warning(msg)
}

func (lg *Logger) Warningf(msg string, v ...interface{}) {
	args := lg.transformParams(v)
	lg.logger.Warningf(msg, args...)
}

func (lg *Logger) transformParams(v []interface{}) []interface{} {
	args := []interface{}{}
	for _, a := range v {
		if reflect.ValueOf(a).Kind() == reflect.Ptr {
			args = append(args, reflect.ValueOf(a).Elem())
		} else {
			args = append(args, a)
		}
	}
	return args
}

func NewLogger(prefix string, _lvl string, console bool) *Logger {
	lg := Logger{}
	lg.console = console
	lg.syslog = false
	lg.prefix = prefix

	lg.configure()

	return &lg
}
