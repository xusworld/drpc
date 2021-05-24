package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// logger level constant
const (
	InfoLogger = iota + 1
	WarningLogger
	DebugLogger
	ErrorLogger
	FatalLogger

	DateTimeFormat = "2006-01-02 15:04:05"
)

var (
	dummyLoggers = make(map[int]*Logger, 5)
)

func init() {

	dummyLoggers[InfoLogger] = &Logger{
		level:   InfoLogger,
		RWMutex: sync.RWMutex{},
	}

	dummyLoggers[WarningLogger] = &Logger{
		level:   WarningLogger,
		RWMutex: sync.RWMutex{},
	}

	dummyLoggers[DebugLogger] = &Logger{
		level:   DebugLogger,
		RWMutex: sync.RWMutex{},
	}

	dummyLoggers[ErrorLogger] = &Logger{
		level:   ErrorLogger,
		RWMutex: sync.RWMutex{},
	}

	dummyLoggers[FatalLogger] = &Logger{
		level:   FatalLogger,
		RWMutex: sync.RWMutex{},
	}
}

func Info(v ...interface{}) {
	dummyLoggers[InfoLogger].Info(v)
}

func Infof(format string, v ...interface{}) {
	dummyLoggers[InfoLogger].Info(format, v)
}

func Warning(v ...interface{}) {
	dummyLoggers[InfoLogger].Warning(v)
}

func Warningf(format string, v ...interface{}) {
	dummyLoggers[InfoLogger].Warning(format, v)
}

func Debug(v ...interface{}) {
	dummyLoggers[DebugLogger].Warning(v)
}

func Debugf(format string, v ...interface{}) {
	dummyLoggers[DebugLogger].Warning(format, v)
}

func Error(v ...interface{}) {
	dummyLoggers[ErrorLogger].Warning(v)
}

func Errorf(format string, v ...interface{}) {
	dummyLoggers[ErrorLogger].Warningf(format, v)
}

func Fatal(v ...interface{}) {
	dummyLoggers[FatalLogger].Warning(v)
}

func Fatalf(format string, v ...interface{}) {
	dummyLoggers[FatalLogger].Warning(format, v)
}

// Logger cool logger
type Logger struct {
	level int
	sync.RWMutex
}

// write
func (l *Logger) write(format string, v ...interface{}) {
	l.RLock()
	defer l.Unlock()

	file, err := os.OpenFile(loggerFileName(l.level), os.O_CREATE|os.O_APPEND, 0664)

	if err != nil {
		// TODO handle this
	}

	pc, filename, line, ok := runtime.Caller(2)
	if !ok {
		// TODO runtime.Caller failed, do something
	}
	record := &Record{
		Timestamp:  time.Now().Format(DateTimeFormat),
		LogLevel:   l.level,
		FileName:   filename,
		LineNumber: line,
		Caller:     runtime.FuncForPC(pc).Name(),
	}

	header := RecordToString(record)
	message := fmt.Sprintf("%s|"+format, header, v)

	_, _ = file.WriteString(message)

	if err := file.Close(); err != nil {
		panic("os.File.Close() error")
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.write("", v)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.write(format, v)
}

func (l *Logger) Warning(v ...interface{}) {
	l.write("", v)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.write(format, v)
}

func (l *Logger) Debug(v ...interface{}) {
	l.write("", v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.write(format, v)
}

func (l *Logger) Error(v ...interface{}) {
	l.write("", v)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.write(format, v)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.write("", v)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.write(format, v)
}

// loggerFileName
func loggerFileName(level int) string {
	return loggerLevelToString(level) + "." + time.Now().Format(DateTimeFormat) + ".log"
}

// loggerLevelToString
func loggerLevelToString(level int) string {
	lookupTable := map[int]string{
		InfoLogger:    "InfoLogger",
		WarningLogger: "WarningLogger",
		DebugLogger:   "DebugLogger",
		ErrorLogger:   "ErrorLogger",
		FatalLogger:   "FatalLogger",
	}

	return lookupTable[level]
}
